package ospath

import "github.com/yyle88/sure"

type existNamespace_Must struct{ T *existNamespace }

func (T *existNamespace) Must() *existNamespace_Must {
	return &existNamespace_Must{T: T}
}
func (T *existNamespace_Must) IsPath(path string) (res bool) {
	res, err1 := T.T.IsPath(path)
	sure.Must(err1)
	return res
}
func (T *existNamespace_Must) IsFile(path string) (res bool) {
	res, err1 := T.T.IsFile(path)
	sure.Must(err1)
	return res
}
func (T *existNamespace_Must) IsRoot(path string) (res bool) {
	res, err1 := T.T.IsRoot(path)
	sure.Must(err1)
	return res
}
func (T *existNamespace_Must) PATH(path string) (res string) {
	res, err1 := T.T.PATH(path)
	sure.Must(err1)
	return res
}
func (T *existNamespace_Must) FILE(path string) (res string) {
	res, err1 := T.T.FILE(path)
	sure.Must(err1)
	return res
}
func (T *existNamespace_Must) ROOT(path string) (res string) {
	res, err1 := T.T.ROOT(path)
	sure.Must(err1)
	return res
}
func (T *existNamespace_Must) MustPath(path string) {
	T.T.MustPath(path)
}
func (T *existNamespace_Must) MustFile(path string) {
	T.T.MustFile(path)
}
func (T *existNamespace_Must) MustRoot(path string) {
	T.T.MustRoot(path)
}

type existNamespace_Soft struct{ T *existNamespace }

func (T *existNamespace) Soft() *existNamespace_Soft {
	return &existNamespace_Soft{T: T}
}
func (T *existNamespace_Soft) IsPath(path string) (res bool) {
	res, err1 := T.T.IsPath(path)
	sure.Soft(err1)
	return res
}
func (T *existNamespace_Soft) IsFile(path string) (res bool) {
	res, err1 := T.T.IsFile(path)
	sure.Soft(err1)
	return res
}
func (T *existNamespace_Soft) IsRoot(path string) (res bool) {
	res, err1 := T.T.IsRoot(path)
	sure.Soft(err1)
	return res
}
func (T *existNamespace_Soft) PATH(path string) (res string) {
	res, err1 := T.T.PATH(path)
	sure.Soft(err1)
	return res
}
func (T *existNamespace_Soft) FILE(path string) (res string) {
	res, err1 := T.T.FILE(path)
	sure.Soft(err1)
	return res
}
func (T *existNamespace_Soft) ROOT(path string) (res string) {
	res, err1 := T.T.ROOT(path)
	sure.Soft(err1)
	return res
}
func (T *existNamespace_Soft) MustPath(path string) {
	T.T.MustPath(path)
}
func (T *existNamespace_Soft) MustFile(path string) {
	T.T.MustFile(path)
}
func (T *existNamespace_Soft) MustRoot(path string) {
	T.T.MustRoot(path)
}

type existNamespace_Omit struct{ T *existNamespace }

func (T *existNamespace) Omit() *existNamespace_Omit {
	return &existNamespace_Omit{T: T}
}
func (T *existNamespace_Omit) IsPath(path string) (res bool) {
	res, err1 := T.T.IsPath(path)
	sure.Omit(err1)
	return res
}
func (T *existNamespace_Omit) IsFile(path string) (res bool) {
	res, err1 := T.T.IsFile(path)
	sure.Omit(err1)
	return res
}
func (T *existNamespace_Omit) IsRoot(path string) (res bool) {
	res, err1 := T.T.IsRoot(path)
	sure.Omit(err1)
	return res
}
func (T *existNamespace_Omit) PATH(path string) (res string) {
	res, err1 := T.T.PATH(path)
	sure.Omit(err1)
	return res
}
func (T *existNamespace_Omit) FILE(path string) (res string) {
	res, err1 := T.T.FILE(path)
	sure.Omit(err1)
	return res
}
func (T *existNamespace_Omit) ROOT(path string) (res string) {
	res, err1 := T.T.ROOT(path)
	sure.Omit(err1)
	return res
}
func (T *existNamespace_Omit) MustPath(path string) {
	T.T.MustPath(path)
}
func (T *existNamespace_Omit) MustFile(path string) {
	T.T.MustFile(path)
}
func (T *existNamespace_Omit) MustRoot(path string) {
	T.T.MustRoot(path)
}
