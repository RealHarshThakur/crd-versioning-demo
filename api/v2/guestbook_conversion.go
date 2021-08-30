package v2

import (
	"fmt"
	"strings"

	v1 "crd-versioning-demo/api/v1"

	"sigs.k8s.io/controller-runtime/pkg/conversion"
)

// ConvertTo is expected to modify its argument to contain the converted object.
// Most of the conversion is straightforward copying, except for converting our changed field.
//ConvertTo converts GuestBook to Hub version(v1)
func (src *Guestbook) ConvertTo(dstRaw conversion.Hub) error {
	dst := dstRaw.(*v1.Guestbook)

	firstName := src.Spec.FirstName
	lastName := src.Spec.LastName

	dst.Spec.FullName = fmt.Sprintf("%s %s", firstName, lastName)
	dst.ObjectMeta = src.ObjectMeta

	return nil
}

/*
ConvertFrom is expected to modify its receiver to contain the converted object.
Most of the conversion is straightforward copying, except for converting our changed field.
*/
// ConvertFrom converts from the Hub version (v1) to this version.
func (dst *Guestbook) ConvertFrom(srcRaw conversion.Hub) error {
	src := srcRaw.(*v1.Guestbook)

	nameParts := strings.Split(src.Spec.FullName, " ")
	dst.Spec.FirstName = nameParts[0]
	dst.Spec.LastName = nameParts[1]
	dst.ObjectMeta = src.ObjectMeta

	return nil
}
