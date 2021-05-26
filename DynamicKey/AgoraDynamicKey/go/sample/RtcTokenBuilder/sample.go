package main

import (
    "fmt"
    rtctokenbuilder "github.com/AgoraIO/Tools/DynamicKey/AgoraDynamicKey/go/src/RtcTokenBuilder"
    "time"
)

func main() {

    appID := "970CA35de60c44645bbae8a215061b33"
    appCertificate := "5CFd2fd1755d40ecb72977518be15d3b"
    channelName := "7d72365eb983485397e3e3f9d460bdda"
    uid := uint32(2882341273)
    uidStr := "2882341273"
    expireTimeInSeconds := uint32(3600)
    currentTimestamp := uint32(time.Now().UTC().Unix())
    expireTimestamp := currentTimestamp + expireTimeInSeconds

    result, err := rtctokenbuilder.BuildTokenWithUID(appID, appCertificate, channelName, uid, rtctokenbuilder.RoleAttendee, expireTimestamp)
    if err != nil {
        fmt.Println(err)
    } else {
        fmt.Printf("Token with uid: %s\n", result)
    }

    result, err = rtctokenbuilder.BuildTokenWithUserAccount(appID, appCertificate, channelName, uidStr, rtctokenbuilder.RoleAttendee, expireTimestamp)
    if err != nil {
        fmt.Println(err)
    } else {
        fmt.Printf("Token with userAccount: %s\n", result)
    }
    result1, err1 := rtctokenbuilder.BuildTokenWithUIDAndPrivilege(appID, appCertificate, channelName, uid, expireTimestamp,
        expireTimestamp, expireTimestamp, expireTimestamp)
    if err1 != nil {
        fmt.Println(err)
    } else {
        fmt.Printf("Token with uid user defined privilege: %s\n", result1)
    }

    result1, err1 = rtctokenbuilder.BuildTokenWithUserAccountAndPrivilege(appID, appCertificate, channelName, uidStr, expireTimestamp,
        expireTimestamp, expireTimestamp, expireTimestamp)

    if err1 != nil {
        fmt.Println(err)
    } else {
        fmt.Printf("Token with userAccount user defined privilege: %s\n", result1)
    }
}
