<h1 align="center">
  mpegts_packets
</h1>

In this example, we'll be using a flowbits Bitstream decoder to decode and display the fields of MPEG transport stream packets.  Information on the MPEG transport stream format can be found [here](https://en.wikipedia.org/wiki/MPEG_transport_stream).

## Instructions

### Download mpegts_packets example

```bash
git clone https://github.com/richinsley/flowbits.git
```

### Compile mpegts_packets example
```bash
cd flowbits/examples/mpegts_packets && go build
```

### Run ffmpeg to create example stream

```bash
ffmpeg -f lavfi -i testsrc=size=320x240:rate=30 -preset ultrafast -vcodec libx264 -tune zerolatency -b 100k -f mpegts "udp://127.0.0.1:1234?pkt_size=188&buffer_size=65535"
```

### Run the mpegts_packets to decode the mpegts packet header
```bash
./mpegts_packets 127.0.0.1:1234
```

### Example output:
```
transport_error_indicator		false
payload_unit_start_indicator		false
transport_priority			false
pid					0x100
transport_scrambling_control		0b0
adaption_field_control			0b11
continuity_counter			0x6
-----Adaption Field:
	discontinuity_indicator			false
	random_access_indicator			false
	elementary_stream_priority_indicator	false
	pcr_flag				false
	opcr_flag				false
	splicing_point_flag			false
	transport_private_data_flag		false
	adaptation_field_extension_flag		false
10001100|11010001|00101110|00001001|01001000|00111011|10010110|10011000
11100111|10111110|01000010|00001000|11010011|01011110|00111001|11101111
11111110|00010001|10011011|11111111|11111111|00100001|00000100|10000100
00011100|11010100|00100110|01111000|01000110|01101100|10011001|10101010
11101110|00101011|11111100|10000100|00010011|01010011|01010101|11110000
10001100|11010010|00010000|01001001|10011001|01001001|00101111|11111111
11101010|10111111|11111101|11111110|00111010|11010001|11111100|00101011
```