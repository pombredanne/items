package SortedMap
type SortedMap struct{
	Basic map[string]interface{}
	ByKey []string
	ByValue []string
	Asc bool
}

//Assume you have done 'sm:=SortedMap{destMap,make([]string,len(destMap)),make([]string,len(destMap),false)}'
//ByValue sorting is only used when its value is  string type
func (sm *SortedMap) Init(){
	if (sm.Asc==true){
		for k,v:=sm.Basic
	}

}