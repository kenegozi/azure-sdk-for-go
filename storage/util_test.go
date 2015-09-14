package storage

import (
	"encoding/xml"
	"io/ioutil"
	"net/url"
	"reflect"
	"strings"
	"time"

	chk "gopkg.in/check.v1"
)

func (s *StorageClientSuite) Test_timeRfc1123Formatted(c *chk.C) {
	now := time.Now().UTC()
	expectedLayout := "Mon, 02 Jan 2006 15:04:05 GMT"
	c.Assert(timeRfc1123Formatted(now), chk.Equals, now.Format(expectedLayout))
}

func (s *StorageClientSuite) Test_mergeParams(c *chk.C) {
	v1 := url.Values{
		"k1": {"v1"},
		"k2": {"v2"}}
	v2 := url.Values{
		"k1": {"v11"},
		"k3": {"v3"}}
	out := mergeParams(v1, v2)
	c.Assert(out.Get("k1"), chk.Equals, "v1")
	c.Assert(out.Get("k2"), chk.Equals, "v2")
	c.Assert(out.Get("k3"), chk.Equals, "v3")
	c.Assert(out["k1"], chk.DeepEquals, []string{"v1", "v11"})
}

func (s *StorageClientSuite) Test_prepareBlockListRequest(c *chk.C) {
	empty := []Block{}
	expected := `<?xml version="1.0" encoding="utf-8"?><BlockList></BlockList>`
	c.Assert(prepareBlockListRequest(empty), chk.DeepEquals, expected)

	blocks := []Block{{"foo", BlockStatusLatest}, {"bar", BlockStatusUncommitted}}
	expected = `<?xml version="1.0" encoding="utf-8"?><BlockList><Latest>foo</Latest><Uncommitted>bar</Uncommitted></BlockList>`
	c.Assert(prepareBlockListRequest(blocks), chk.DeepEquals, expected)
}

func (s *StorageClientSuite) Test_xmlUnmarshal(c *chk.C) {
	xml := `<?xml version="1.0" encoding="utf-8"?>
	<Blob>
		<Name>myblob</Name>
	</Blob>`
	var blob Blob
	body := ioutil.NopCloser(strings.NewReader(xml))
	c.Assert(xmlUnmarshal(body, &blob), chk.IsNil)
	c.Assert(blob.Name, chk.Equals, "myblob")
}

func (s *StorageClientSuite) Test_xmlMarshal(c *chk.C) {
	type t struct {
		XMLName xml.Name `xml:"S"`
		Name    string   `xml:"Name"`
	}

	b := t{Name: "myblob"}
	expected := `<S><Name>myblob</Name></S>`
	r, i, err := xmlMarshal(b)
	c.Assert(err, chk.IsNil)
	o, err := ioutil.ReadAll(r)
	c.Assert(err, chk.IsNil)
	out := string(o)
	c.Assert(out, chk.Equals, expected)
	c.Assert(i, chk.Equals, len(expected))
}

func init() {
	initHeaderPropertiesMapsForType(reflect.TypeOf(&someProperties{}))
}

type someProperties struct {
	Number              int64             `header:"num"`
	NumberPtr           *int64            `header:"numptr"`
	NumberPtrNil        *int64            `header:"numptrnil"`
	Text                string            `header:"str"`
	MixedCase           string            `header:"miXed-Case"`
	EmptyText           string            `header:"emptystring"`
	Ignored             string            //no header tag
	MultiHeader         string            `header:"h1,h2"`
	My                  myString          `header:"my"`
	MultiHeaderPrefixed string            `header:"prefix-h3,prefix-h4"`
	Untyped             map[string]string `header:"@"`
	UntypedPrefixed     map[string]string `header:"@prefix-"`
}

type propsThatIsAMAp map[string]string

type myString string

func (s *StorageBlobSuite) Test_marshalProperties(c *chk.C) {
	props := someProperties{}
	props.Number = 1
	n := int64(2)
	props.NumberPtr = &n
	props.Text = "Hello"
	props.Ignored = "Whatever"
	props.MultiHeader = "Appear twice"
	props.My = "str"
	props.MixedCase = "mixed"
	props.MultiHeaderPrefixed = "maybe overriden"
	props.UntypedPrefixed = map[string]string{"h3": "prefixed", "h5": "hidden", "h6": "stay put"}
	props.Untyped = map[string]string{"prefix-h5": "overriden", "foo": "bar"}

	hds := map[string]string{}
	err := marshalProperties(&props, hds)
	c.Assert(err, chk.IsNil)

	c.Assert(hds["num"], chk.Equals, "1")
	c.Assert(hds["numptr"], chk.Equals, "2")
	c.Assert(hds["str"], chk.Equals, "Hello")
	c.Assert(hds["h1"], chk.Equals, "Appear twice")
	c.Assert(hds["h2"], chk.Equals, "Appear twice")
	c.Assert(hds["my"], chk.Equals, "str")
	c.Assert(hds["miXed-Case"], chk.Equals, "mixed")
	c.Assert(hds["prefix-h3"], chk.Equals, "prefixed")
	c.Assert(hds["prefix-h4"], chk.Equals, "maybe overriden")
	c.Assert(hds["prefix-h5"], chk.Equals, "overriden")
	c.Assert(hds["prefix-h6"], chk.Equals, "stay put")
	c.Assert(hds["foo"], chk.Equals, "bar")
	c.Assert(len(hds), chk.Equals, 12)
}

func (s *StorageBlobSuite) Test_unmarshalProperties(c *chk.C) {
	hds := map[string]string{
		"num":        "1",
		"numptr":     "2",
		"str":        "Hello",
		"h1":         "Appear twice",
		"h2":         "Appear twice",
		"my":         "str",
		"Mixed-Case": "mixed",
		"prefix-a":   "prefixedv1",
		"prefix-b":   "prefixedv2",
		"foo":        "fooval",
		"bar":        "barval",
	}

	props := someProperties{}
	err := unmarshalProperties(hds, &props)
	c.Assert(err, chk.IsNil)

	c.Assert(props.Number, chk.Equals, int64(1))
	c.Assert(*props.NumberPtr, chk.Equals, int64(2))
	c.Assert(props.NumberPtrNil, chk.IsNil)
	c.Assert(props.Text, chk.Equals, "Hello")
	c.Assert(props.MultiHeader, chk.Equals, "Appear twice")
	c.Assert(props.My, chk.Equals, myString("str"))
	c.Assert(props.MixedCase, chk.Equals, "mixed")

	c.Assert(len(props.UntypedPrefixed), chk.Equals, 2)
	c.Assert(props.UntypedPrefixed["a"], chk.Equals, "prefixedv1")
	c.Assert(props.UntypedPrefixed["b"], chk.Equals, "prefixedv2")

	c.Assert(len(props.Untyped), chk.Equals, 2)
	c.Assert(props.Untyped["foo"], chk.Equals, "fooval")
	c.Assert(props.Untyped["bar"], chk.Equals, "barval")
}

func (s *StorageBlobSuite) Test_marshalProperties_nonRegisteredType_ErrorNoPanic(c *chk.C) {
	err := marshalProperties(&Blob{}, map[string]string{})
	c.Assert(err, chk.NotNil)
}

func (s *StorageBlobSuite) Test_unmarshalProperties_nonRegisteredType_ErrorNoPanic(c *chk.C) {
	err := unmarshalProperties(map[string]string{}, &Blob{})
	c.Assert(err, chk.NotNil)
}
