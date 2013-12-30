package dictionary

import "fmt"

type item struct {
	data int
	key string
}
type dictionary struct {
	Count int
	keys []string
	items []*item
}

func (d *dictionary) Insert(key string, data int) {
	if d.Contains(key) == false {
		i := initItem(key, data)
		d.keys = append(d.keys, key)
		d.items = append(d.items, i)
		d.Count = d.Count + 1
	}
}

func (d *dictionary) Remove(key string) (int) {
	if d.Count == 0 {
		panic("Dictionary is empty")
	}
	for i:=0;i<d.Count;i++ {
		if(d.items[i].key == key) {
			data := d.items[i].data
			d.items = append(d.items[:i], d.items[i+1:]...)
			d.keys = append(d.keys[:i], d.keys[i+1:]...)
			d.Count = d.Count - 1
			return data
		}
	}
	panic(sout("Key \"%s\" not found", key))
}

func (d *dictionary) Contains(key string) (bool) {
	if d.Count == 0 {
		return false
	}
	for _,v := range d.keys {
		if v == key {
			return true
		}
	}
	return false
}

func (d *dictionary) Search(key string) (int) {
	if d.Count == 0 {
		panic("Dictionary is empty")
	}
	for _,v := range d.items {
		if v.key == key {
			return v.data
		}
	}
	panic(sout("Key \"%s\" not found", key))
}

func (d *dictionary) PrintKeys() {
	if d.Count == 0 {
		out("dictionary is empty\n")
		return
	}
	for _,v := range d.keys {
		out("%s ", v)
	}
	out("\n")
}

func InitDictionary() (*dictionary) {
	d := new(dictionary)
	d.Count = 0
	d.keys = make([]string, 0)
	//d.items = make([]item, 0)
	return d
}

func initItem(key string, data int) (*item) {
	i := new(item)
	i.data = data
	i.key = key
	return i
}

// Shorthand function for fmt.Printf
func out(format string, v...interface{}) {
	fmt.Printf(format, v...)
}
func sout(format string, v...interface{}) (string) {
	return fmt.Sprintf(format, v...)
}