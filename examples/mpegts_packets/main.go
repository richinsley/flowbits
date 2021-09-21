package main

import (
	"bytes"
	"fmt"
	"net"
	"os"
	"strconv"

	"github.com/richinsley/flowbits"
	"github.com/richinsley/purtybits"
)

type MPEGTS_Header struct {
	transport_error_indicator    bool
	payload_unit_start_indicator bool
	transport_priority           bool
	pid                          uint16
	transport_scrambling_control uint8
	adaption_field_control       uint8
	continuity_counter           uint8
	adaption_field               *MPEGTS_Header_AdaptionField
	payload                      *[]uint8
}

type MPEGTS_Header_AdaptionField struct {
	adaption_field_length                uint8
	discontinuity_indicator              bool
	random_access_indicator              bool
	elementary_stream_priority_indicator bool
	pcr_flag                             bool
	opcr_flag                            bool
	splicing_point_flag                  bool
	transport_private_data_flag          bool
	adaptation_field_extension_flag      bool
}

// we'll use purty bits to format our binary payload output
var purty *purtybits.PurtyBits = nil

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stderr, "Usage:%s host:port", os.Args[0])
		os.Exit(1)
	}

	// clear the terminal
	fmt.Print("\u001b[2J")

	service := os.Args[1]
	udpAddr, err := net.ResolveUDPAddr("udp", service)
	checkError(err)

	conn, err := net.ListenUDP("udp", udpAddr)
	checkError(err)
	conn.SetReadBuffer(188 * 100)

	// configure purtybits to output just the binary data
	purty = purtybits.NewPurtyBits(8, purtybits.HexCodeNone)

	for {
		handleClient(conn)
	}
}

// GenerateOutput formats the header and it's payload into an
// array of strings for output to the terminal
func GenerateOutput(header *MPEGTS_Header) []string {
	var retv []string
	retv = append(retv, "transport_error_indicator\t\t"+strconv.FormatBool(header.transport_error_indicator)+"\u001b[0K")
	retv = append(retv, "payload_unit_start_indicator\t\t"+strconv.FormatBool(header.payload_unit_start_indicator)+"\u001b[0K")
	retv = append(retv, "transport_priority\t\t\t"+strconv.FormatBool(header.transport_priority)+"\u001b[0K")
	retv = append(retv, "pid\t\t\t\t\t0x"+strconv.FormatUint(uint64(header.pid), 16)+"\u001b[0K")
	retv = append(retv, "transport_scrambling_control\t\t0b"+strconv.FormatUint(uint64(header.transport_scrambling_control), 2)+"\u001b[0K")
	retv = append(retv, "adaption_field_control\t\t\t0b"+strconv.FormatUint(uint64(header.adaption_field_control), 2)+"\u001b[0K")
	retv = append(retv, "continuity_counter\t\t\t0x"+strconv.FormatUint(uint64(header.continuity_counter), 16)+"\u001b[0K")

	if header.adaption_field != nil {
		retv = append(retv, "-----Adaption Field:"+"\u001b[0K")
		retv = append(retv, "\tdiscontinuity_indicator\t\t\t"+strconv.FormatBool(header.adaption_field.discontinuity_indicator)+"\u001b[0K")
		retv = append(retv, "\trandom_access_indicator\t\t\t"+strconv.FormatBool(header.adaption_field.random_access_indicator)+"\u001b[0K")
		retv = append(retv, "\telementary_stream_priority_indicator\t"+strconv.FormatBool(header.adaption_field.elementary_stream_priority_indicator)+"\u001b[0K")
		retv = append(retv, "\tpcr_flag\t\t\t\t"+strconv.FormatBool(header.adaption_field.pcr_flag)+"\u001b[0K")
		retv = append(retv, "\topcr_flag\t\t\t\t"+strconv.FormatBool(header.adaption_field.opcr_flag)+"\u001b[0K")
		retv = append(retv, "\tsplicing_point_flag\t\t\t"+strconv.FormatBool(header.adaption_field.splicing_point_flag)+"\u001b[0K")
		retv = append(retv, "\ttransport_private_data_flag\t\t"+strconv.FormatBool(header.adaption_field.transport_private_data_flag)+"\u001b[0K")
		retv = append(retv, "\tadaptation_field_extension_flag\t\t"+strconv.FormatBool(header.adaption_field.adaptation_field_extension_flag)+"\u001b[0K")
	} else {
		retv = append(retv, "-----Adaption Field: [None]"+"\u001b[0K")
		retv = append(retv, "\tdiscontinuity_indicator\t\t\t--\u001b[0K")
		retv = append(retv, "\trandom_access_indicator\t\t\t--\u001b[0K")
		retv = append(retv, "\telementary_stream_priority_indicator\t--\u001b[0K")
		retv = append(retv, "\tpcr_flag\t\t\t\t--\u001b[0K")
		retv = append(retv, "\topcr_flag\t\t\t\t--\u001b[0K")
		retv = append(retv, "\tsplicing_point_flag\t\t\t--\u001b[0K")
		retv = append(retv, "\ttransport_private_data_flag\t\t--\u001b[0K")
		retv = append(retv, "\tadaptation_field_extension_flag\t\t--\u001b[0K")
	}

	if header.payload != nil {
		// create colorized binary output for the payload
		retv = append(retv, purty.BufferToStrings(*header.payload)...)
	}

	return retv
}

// GetMPEGTSAdaptionField parses the adaption field of an mpegts header
func GetMPEGTSAdaptionField(foor *flowbits.Bitstream) *MPEGTS_Header_AdaptionField {
	adaption_field := &MPEGTS_Header_AdaptionField{}
	adaption_field.adaption_field_length, _ = foor.GetUint8()
	if adaption_field.adaption_field_length == 0 {
		// odd, but possible
		return adaption_field
	}

	adaption_field.discontinuity_indicator, _ = foor.GetBool()
	adaption_field.random_access_indicator, _ = foor.GetBool()
	adaption_field.elementary_stream_priority_indicator, _ = foor.GetBool()
	adaption_field.pcr_flag, _ = foor.GetBool()
	adaption_field.opcr_flag, _ = foor.GetBool()
	adaption_field.splicing_point_flag, _ = foor.GetBool()
	adaption_field.transport_private_data_flag, _ = foor.GetBool()
	adaption_field.adaptation_field_extension_flag, _ = foor.GetBool()

	// The adaption_field_length is the number of bytes in the adaptation field, not including the adaption_field_length byte.
	// We're not going to parse the remainder of the adaption field, so we'll just skip it.  We've already read one byte of flags above.
	foor.Skipbits(uint32(adaption_field.adaption_field_length-1) * 8)

	return adaption_field
}

// GetMPEGTSHeader parses an mpegts header
func GetMPEGTSHeader(foor *flowbits.Bitstream) *MPEGTS_Header {
	mpegts_header := &MPEGTS_Header{}

	// populate the fields of the mpegts header from the bitstream
	mpegts_header.transport_error_indicator, _ = foor.GetBool()
	mpegts_header.payload_unit_start_indicator, _ = foor.GetBool()
	mpegts_header.transport_priority, _ = foor.GetBool()
	mpegts_header.pid, _ = foor.GetWithBitCountBigUint16(13)
	mpegts_header.transport_scrambling_control, _ = foor.GetWithBitCountUint8(2)
	mpegts_header.adaption_field_control, _ = foor.GetWithBitCountUint8(2)
	mpegts_header.continuity_counter, _ = foor.GetWithBitCountUint8(4)

	// is there an adaption field
	if mpegts_header.adaption_field_control&0x02 != 0 {
		mpegts_header.adaption_field = GetMPEGTSAdaptionField(foor)
	}

	// anything from this point is payload if adaption_field_control&1==1
	if mpegts_header.adaption_field_control&0x01 != 0 {
		payload_length := uint64(188 - (foor.GetPos() / 8))
		payload := make([]uint8, payload_length)
		mpegts_header.payload = &payload
		bytes_read, err := foor.GetBuffer(payload, payload_length)
		if bytes_read != payload_length || err != nil {
			fmt.Println("nope")
		}
	}

	return mpegts_header
}

func handleClient(conn *net.UDPConn) {
	// mpegts packets are (almost) always 188 bytes in length
	var buf [188]byte

	_, _, err := conn.ReadFrom(buf[0:])
	if err != nil {
		return
	}

	// create a byte reader from the buffer and instantiate a flowbits Bitstream decoder with it
	r := bytes.NewReader(buf[:])
	foor := flowbits.NewBitstreamDecoder(r, 188)

	// all mpegts header fields are in Big Endian format and begin with the byte 0x47
	sync_byte, _ := foor.GetUint8()
	if sync_byte != 0x47 {
		fmt.Println("Incorrect sync byte")
		return
	}

	// parse the header and it's payload
	mpegts_header := GetMPEGTSHeader(foor)

	// format the header's fields and payload into a readable format
	tty_lines := GenerateOutput(mpegts_header)

	// re-home the terminal and display the header
	fmt.Print("\033[H")
	for i := 0; i < len(tty_lines); i++ {
		fmt.Println(tty_lines[i])
	}

	// clear the terminal from the current cursor to the the end of the screen
	fmt.Println("\u001b[0J")
}

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error %s", err.Error())
		os.Exit(1)
	}
}
