package component

func RemoveEmpty(a []string) (ret []string) {
	for _, v := range a {
		if v == "" {
			continue
		}
		ret = append(ret, v)
	}
	return
}
