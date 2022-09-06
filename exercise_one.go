package go_unit_test_bootcamp

func FindMissingDrone(droneIds []int) int {
	// TODO: implement this
	entry := make(map[int]int)
	for i:=0;i<len(droneIds);i++{
		if _, ok := entry[droneIds[i]]; ok {
			entry[droneIds[i]]++
		}else{
			entry[droneIds[i]]=1
		}		
	}
	for k,v := range entry{
		if v == 1  {
			return k
		}
	}
	return -1
}
