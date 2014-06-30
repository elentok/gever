package gever

func (v *Version) CreateRC() {
	rc := v.RC() + 1
	v.SetRC(rc)
}

func (v *Version) CreateHotfix() {
	hotfix := v.Hotfix() + 1
	v.SetHotfix(hotfix)
}

func (v *Version) CreatePatch() {
	patch := v.Patch() + 1
	v.SetPatch(patch)
	v.SetSpecial("")
}

func (v *Version) CreateMinor() {
	minor := v.Minor() + 1
	v.SetMinor(minor)
	v.SetPatch(0)
	v.SetSpecial("")
}

func (v *Version) CreateMajor() {
	major := v.Major() + 1
	v.SetMajor(major)
	v.SetMinor(0)
	v.SetPatch(0)
	v.SetSpecial("")
}
