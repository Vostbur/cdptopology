package tools

import (
	"encoding/json"
	"io/ioutil"
)

type linksStruct struct {
	Id        int    `json:"id"`
	Source    int    `json:"source"`
	SrcDevice string `json:"srcDevice"`
	SrcIfName string `json:"srcIfName"`
	Target    int    `json:"target"`
	TgtDevice string `json:"tgtDevice"`
	TgtIfName string `json:"tgtIfName"`
}
type nodesStruct struct {
	Icon string `json:"icon"`
	Id   int    `json:"id"`
	Name string `json:"name"`
}
type linksStructJson struct {
	Links []linksStruct `json:"links"`
}
type nodesStructJson struct {
	Nodes []nodesStruct `json:"nodes"`
}

// test
func BuildJson(devs []Device, devMap map[string]int) error {
	var ls, lsTmp linksStructJson
	var ns nodesStructJson
	var i int
	for index, device := range devs {
		ns.Nodes = append(ns.Nodes, nodesStruct{
			Icon: "router",
			Id:   index,
			Name: device.Hostname,
		})
		for _, neighboors := range device.Neighboors {
			lsTmp.Links = append(lsTmp.Links, linksStruct{
				Id:        i,
				Source:    index,
				SrcDevice: device.Hostname,
				SrcIfName: neighboors.LocalInterface,
				Target:    devMap[neighboors.Hostname],
				TgtDevice: neighboors.Hostname,
				TgtIfName: neighboors.RemoteInterface,
			})
			i++
		}
	}
	// del dupl
	var flag bool
	var newIndLink int
	ls.Links = append(ls.Links, lsTmp.Links[0])
	for _, linkTmp := range lsTmp.Links[1:] {
		flag = true
		for _, j := range ls.Links {
			if j.Source == linkTmp.Target && j.Target == linkTmp.Source {
				flag = false
				break
			}
		}
		if !flag {
			continue
		}
		newIndLink++
		linkTmp.Id = newIndLink
		ls.Links = append(ls.Links, linkTmp)
	}
	lsJson, err := json.MarshalIndent(ls, "", "\t")
	if err != nil {
		return err
	}
	nsJson, err := json.MarshalIndent(ns, "", "\t")
	if err != nil {
		return err
	}
	strLsJon := string(lsJson)
	strLsJon = "var topologyData = " + strLsJon[:len(strLsJon)-2] + ","
	strNsJson := string(nsJson)
	strNsJson = strNsJson[1:] + ";\n"
	resultJson := strLsJon + strNsJson
	if err = ioutil.WriteFile("topology.js", []byte(resultJson), 0644); err != nil {
		return err
	}
	return nil
}
