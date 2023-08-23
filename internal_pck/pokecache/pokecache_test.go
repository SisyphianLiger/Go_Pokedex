package pokecache


import "testing"


func TestCreateCache( t *testing.T) {
    cache := NewCache() 
    if cache.cache == nil {
        t.Error("cache is nil")
    }
}




func TestAddCache( t *testing.T) {
    cache := NewCache() 

    cases := []struct {
        inputKey string 
        inputVal []byte 
    }{
        {
            inputKey: "key1",
            inputVal: []byte("val1"),
        },
        {
            inputKey: "key2",
            inputVal: []byte("val2"),
        },
        {
            inputKey: "",
            inputVal: []byte("val3"),
        },
    }

    
    for _ , cas := range cases {
        cache.Add(cas.inputKey, cas.inputVal)
        actual, ok := cache.Get(cas.inputKey)
        
        if !ok {
            t.Errorf("%s is not found", cas.inputKey)
            continue
        }
        if string(actual) != string(cas.inputVal) {
            t.Errorf("%s val does not match %s", string(actual) ,string(cas.inputVal))
        }
        continue
    }
}


