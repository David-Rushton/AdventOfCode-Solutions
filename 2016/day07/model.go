package main

type ip7Address struct {
	sections []ip7Section
}

func (ip *ip7Address) supportsTLS() bool {
	var abbaFound bool
	for _, section := range ip.sections {
		if section.containsAbba() {
			if section.isHypernetSequence {
				return false
			}

			abbaFound = true
		}
	}

	return abbaFound
}

func (ip *ip7Address) supportsSSL() bool {
	abas := []string{}
	babs := []string{}

	for _, section := range ip.sections {
		abas = append(abas, section.getAreaBroadcastAccessors()...)
		babs = append(babs, section.getByteAllocationBlock()...)
	}

	for _, aba := range abas {
		for _, bab := range babs {
			if aba[0] == bab[1] && bab[0] == aba[1] {
				return true
			}
		}
	}

	return false
}

func (ip ip7Address) String() string {
	var result string

	for _, section := range ip.sections {
		if section.isHypernetSequence {
			result += "[" + section.value + "]"
		} else {
			result += section.value
		}
	}

	return result
}

type ip7Section struct {
	value              string
	isHypernetSequence bool
}

func (ip *ip7Section) containsAbba() bool {
	if len(ip.value) >= 4 {
		// ABBA when any 4 consequence characters form a palindrome, & contain 2 different characters.
		// - example: ABBA ✅
		// - example: ABCD ❌
		// - example: XXXX ❌
		// - example: DAAD ✅
		for i := 3; i < len(ip.value); i++ {
			if ip.value[i-3] == ip.value[i] {
				if ip.value[i-2] == ip.value[i-1] {
					if ip.value[i-1] != ip.value[i] {
						return true
					}
				}
			}
		}
	}

	return false
}

func (ip *ip7Section) getAreaBroadcastAccessors() []string {
	result := []string{}

	if !ip.isHypernetSequence {
		for i := 2; i < len(ip.value); i++ {
			if ip.value[i-2] == ip.value[i] && ip.value[i-1] != ip.value[i] {
				result = append(result, ip.value[i-2:i+1])
			}
		}
	}

	return result
}

func (ip *ip7Section) getByteAllocationBlock() []string {
	result := []string{}

	if ip.isHypernetSequence {
		for i := 2; i < len(ip.value); i++ {
			if ip.value[i-2] == ip.value[i] && ip.value[i-1] != ip.value[i] {
				result = append(result, ip.value[i-2:i+1])
			}
		}
	}

	return result
}
