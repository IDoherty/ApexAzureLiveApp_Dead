package aggFuncs

import (
        "log"       
		"fmt"
		//"time"
		"encoding/hex"
		"encoding/binary"
)

func procPackets(inUDPChan <-chan string, outUDPChan chan<- string){

	// Build Variables
	var arrayLastPackets [50]LastPacket
	//valPkt:= 0;
	
	//fmt.Println("9")
	
	keepAlivePacketID := []byte("55DD")
	packetTest := binary.LittleEndian.Uint16(keepAlivePacketID)
	
	fmt.Println(packetTest)
	//fmt.Println()
	
	for{
		returnedData := <-inUDPChan
		
		// Revert Data to []byte
		destringifiedData, err := hex.DecodeString(returnedData)
		if err != nil {
			log.Fatal(err)
		}
		
		fmt.Printf("%s", hex.Dump(destringifiedData))
		fmt.Println()
		
		//fmt.Println("10")
		
		// Slice Packet Length
		//packetLen := binary.LittleEndian.Uint16(destringifiedData[2:4])
		//fmt.Println(packetLen)
		
		// Filter out KA responses and test Validity for incoming packets
		if binary.LittleEndian.Uint16(destringifiedData[0:2]) == packetTest{
			//fmt.Println("Keep Alive Packet", )
			//fmt.Println()

		}else{
			seqNo 	:= destringifiedData[4:5]
			slotNo 	:= destringifiedData[5:6]
			gpsTime	:= binary.BigEndian.Uint32(destringifiedData[8:12])
			
			fmt.Println(seqNo)
			fmt.Println(slotNo)
			fmt.Println(gpsTime)
			
			valid, pktType := TestValidity(&arrayLastPackets[slotNo[0]], seqNo[0], gpsTime)
			
			//countChan <- pktType
			pktType++
			
			if valid{
				//valPkt++
				outUDPChan <- returnedData
			}
			
		}
		
	}

}
