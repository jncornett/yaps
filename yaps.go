package main

import (
    "fmt"
    "os"
    // "github.com/fatih/color"
    "github.com/google/gopacket"
    // "github.com/google/gopacket/layers"
    "github.com/google/gopacket/pcap"
)

const (
    SNAPLEN = 1600
    MAGIC = 26
)

func nextAlias(a map[string]string) string {
    bytes := []byte{}
    ctr := len(a)

    for ctr > 0 {
        div := ctr / MAGIC
        ctr = ctr % MAGIC
        bytes = append(bytes, 'a' + byte(div))
    }

    return string(bytes)
}

func handlePacket(a map[string]string, p gopacket.Packet) {
    // red := color.New(color.FgRed).SprintFunc()
    // blue := color.New(color.FgBlue).SprintFunc()

    // if layer := p.Layer(layers.LayerTypeEthernet); layer != nil {
    //     eth, _ := layer.(*layers.Ethernet)
    //     fmt.Println(red(eth.SrcMAC.String()), ">", blue(eth.DstMAC.String()))
    // }

    // if layer := p.Layer(layers.LayerTypeIPv4); layer != nil {
    //     var src, dst string
    //     var ok bool

    //     ip4, _ := layer.(*layers.IPv4)

    //     if src, ok = a[ip4.SrcIP.String()]; !ok {
    //         src = nextAlias(a)
    //         a[ip4.SrcIP.String()] = src
    //     }

    //     if dst, ok := a[ip4.DstIP.String()]; !ok {
    //         dst = nextAlias(a)
    //         a[ip4.DstIP.String()] = dst
    //     }

    //     fmt.Println(red(src), ">", blue(dst))
    // } else if layer := p.Layer(layers.LayerTypeIPv6); layer != nil {
    //     var src, dst string
    //     var ok bool

    //     ip6, _ := layer.(*layers.IPv6)

    //     if src, ok = a[ip6.SrcIP.String()]; !ok {
    //         src = nextAlias(a)
    //         a[ip6.SrcIP.String()] = src
    //     }

    //     if dst, ok := a[ip6.DstIP.String()]; !ok {
    //         dst = nextAlias(a)
    //         a[ip6.DstIP.String()] = dst
    //     }

    //     fmt.Println(red(src), ">", blue(dst))
    // }

    // for _, layer := range p.Layers() {
    //     fmt.Println(blue(layer.LayerType()))
    //     fmt.Println(layer)
    // }

    fmt.Println(p)
}

func main() {
    a := make(map[string]string)
    dev := os.Args[1]

    handle, err := pcap.OpenLive(dev, SNAPLEN, true, pcap.BlockForever)
    if err != nil { panic(err) }

    packetSource := gopacket.NewPacketSource(handle, handle.LinkType())
    for p := range packetSource.Packets() {
        handlePacket(a, p)
    }
}
