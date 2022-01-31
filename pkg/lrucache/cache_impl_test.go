package lrucache_test

import (
	"github.com/asavt7/lrucache/pkg/lrucache"
	"reflect"
	"strconv"
	"testing"
)

func TestNewLRUCache(t *testing.T) {

	t.Run("invalid args", func(t *testing.T) {
		defer func() {
			err := recover()
			if err == nil {
				t.Errorf("expected panic")
			}
		}()
		var _ lrucache.LRUCache = lrucache.NewLRUCache(0)

	})
}

func TestInMemoryLRUCache_Get(t *testing.T) {
	t.Run("empty cache and no such keys", func(t *testing.T) {
		var cache lrucache.LRUCache = lrucache.NewLRUCache(10)
		for i := 0; i < 10; i++ {
			v, ok := cache.Get(strconv.Itoa(i))
			if v != "" || ok {
				t.Errorf("empty map should return empty and !ok")
			}
		}
	})

	t.Run("not empty cache and no such keys", func(t *testing.T) {
		var cache lrucache.LRUCache = lrucache.NewLRUCache(10)
		cache.Add("dasda", "99")
		cache.Add("jgda", "299")
		cache.Add("das", "199")

		for i := 0; i < 10; i++ {
			v, ok := cache.Get(strconv.Itoa(i))
			if v != "" || ok {
				t.Errorf("empty map should return empty and !ok")
			}
		}
	})

	t.Run("not empty cache and no such keys", func(t *testing.T) {
		var cache lrucache.LRUCache = lrucache.NewLRUCache(10)
		for i := 0; i < 3; i++ {
			k := strconv.Itoa(i)
			cache.Add(k, k)
		}

		for i := 0; i < 3; i++ {
			k := strconv.Itoa(i)
			v, ok := cache.Get(strconv.Itoa(i))
			if v != k || !ok {
				t.Errorf("unexpected value")
			}
		}
	})

	t.Run("cache full", func(t *testing.T) {
		var cache lrucache.LRUCache = lrucache.NewLRUCache(4)
		for i := 0; i < 5; i++ {
			k := strconv.Itoa(i)
			cache.Add(k, k)
		}

		v, ok := cache.Get("0")
		if v != "" || ok {
			t.Errorf("it was expected that there would be no %s key", "0")
		}

		for i := 1; i < 5; i++ {
			k := strconv.Itoa(i)
			v, ok := cache.Get(k)
			if v != k || !ok {
				t.Errorf("unexpected value")
			}
			v, ok = cache.Get(k)
			if v != k || !ok {
				t.Errorf("unexpected value")
			}
		}
	})

	t.Run("cache full rewrite", func(t *testing.T) {
		var cache lrucache.LRUCache = lrucache.NewLRUCache(4)
		for i := 0; i < 10; i++ {
			k := strconv.Itoa(i)
			cache.Add(k, k)
		}

		for i := 0; i < 6; i++ {
			k := strconv.Itoa(i)
			v, ok := cache.Get(k)
			if v != "" || ok {
				t.Errorf("it was expected that there would be no %s key", "0")
			}
		}

		for i := 6; i < 10; i++ {
			k := strconv.Itoa(i)
			v, ok := cache.Get(k)
			if v != k || !ok {
				t.Errorf("unexpected value")
			}
		}
	})
}

func TestInMemoryLRUCache_Remove(t *testing.T) {
	t.Run("empty cache", func(t *testing.T) {
		var cache lrucache.LRUCache = lrucache.NewLRUCache(10)
		for i := 0; i < 3; i++ {
			k := strconv.Itoa(i)
			ok := cache.Remove(k)
			if ok {
				t.Errorf("expected return !ok for empty cache")
			}
		}

		v, ok := cache.Get("0")
		if v != "" || ok {
			t.Errorf("expected there is no key in cache")
		}
	})

	t.Run("one element", func(t *testing.T) {
		var cache lrucache.LRUCache = lrucache.NewLRUCache(10)
		cache.Add("0", "0")
		cache.Remove("0")
		for i := 0; i < 3; i++ {
			k := strconv.Itoa(i)
			ok := cache.Remove(k)
			if ok {
				t.Errorf("expected return !ok for empty cache")
			}
		}

	})

	t.Run("not empty cache", func(t *testing.T) {
		var cache lrucache.LRUCache = lrucache.NewLRUCache(10)
		for i := 0; i < 10; i++ {
			k := strconv.Itoa(i)
			cache.Add(k, k)
		}

		for i := 0; i < 10; i++ {
			k := strconv.Itoa(i)
			ok := cache.Remove(k)
			if !ok {
				t.Errorf("expected return ok for key, that exists in cache")
			}
		}

		for i := 0; i < 10; i++ {
			k := strconv.Itoa(i)
			v, ok := cache.Get(k)
			if v != "" || ok {
				t.Errorf("expected there is no key in cache")
			}
		}

		for i := 0; i < 10; i++ {
			k := strconv.Itoa(i)
			cache.Add(k, k)
		}
		for i := 0; i < 10; i++ {
			k := strconv.Itoa(i)
			v, ok := cache.Get(k)
			if v != k || !ok {
				t.Errorf("expected key is in cache")
			}
		}
	})
}

func TestInMemoryLRUCache_Keys(t *testing.T) {
	t.Run("keys", func(t *testing.T) {
		var cache lrucache.LRUCache = lrucache.NewLRUCache(10)

		var expectedKeys []string
		for i := 0; i < 10; i++ {
			k := strconv.Itoa(i)
			cache.Add(k, k)
			expectedKeys = append([]string{k}, expectedKeys...)
		}

		keys := cache.(*lrucache.InMemoryLRUCache).Keys()

		if !reflect.DeepEqual(keys, expectedKeys) {
			t.Errorf("expected keys %+v != actual  %+v", expectedKeys, keys)
		}
	})

}

func TestInMemoryLRUCache_Add(t *testing.T) {
	t.Run("rewrite keys", func(t *testing.T) {
		var cache lrucache.LRUCache = lrucache.NewLRUCache(10)

		for i := 0; i < 100; i++ {
			k := strconv.Itoa(i % 10)
			v := strconv.Itoa(i)
			cache.Add(k, v)
		}

		for i := 90; i < 100; i++ {
			k := strconv.Itoa(i % 10)
			expectedVal := strconv.Itoa(i)
			v, ok := cache.Get(k)
			if !ok || expectedVal != v {
				t.Errorf("expected %s got %s for key %s", expectedVal, v, k)
			}
		}
	})

	t.Run("size = 1", func(t *testing.T) {
		var cache lrucache.LRUCache = lrucache.NewLRUCache(1)
		for i := 0; i < 100; i++ {
			k := strconv.Itoa(i % 10)
			v := strconv.Itoa(i)
			cache.Add(k, v)
		}
		for i := 90; i < 99; i++ {
			k := strconv.Itoa(i % 10)
			v, ok := cache.Get(k)
			if ok || "" != v {
				t.Errorf("expected '' got %s for key %s", v, k)
			}
		}
	})

	t.Run("rm last added elem", func(t *testing.T) {
		var cache lrucache.LRUCache = lrucache.NewLRUCache(10)
		for i := 0; i < 10; i++ {
			k := strconv.Itoa(i)
			v := strconv.Itoa(i)
			cache.Add(k, v)
		}
		if ok := cache.Remove("9"); !ok {
			t.Errorf("expected key in cache is removed")
		}
	})

	t.Run("rm mid-added elem", func(t *testing.T) {
		var cache lrucache.LRUCache = lrucache.NewLRUCache(10)
		for i := 0; i < 10; i++ {
			k := strconv.Itoa(i)
			v := strconv.Itoa(i)
			cache.Add(k, v)
		}
		if ok := cache.Remove("5"); !ok {
			t.Errorf("expected key in cache is removed")
		}
		if v, ok := cache.Get("5"); v != "" || ok {
			t.Errorf("expected cache not contains key")
		}
	})

	t.Run("size = 2", func(t *testing.T) {
		var cache lrucache.LRUCache = lrucache.NewLRUCache(2)
		cache.Add("0", "0")
		cache.Add("1", "1")
		cache.Remove("1")
		cache.Remove("0")

		if v, ok := cache.Get("0"); v != "" || ok {
			t.Errorf("expected empty cache")
		}
		if v, ok := cache.Get("1"); v != "" || ok {
			t.Errorf("expected empty cache")
		}

		cache.Add("1", "1")
		cache.Add("0", "0")
		cache.Remove("1")
		cache.Remove("0")

		if v, ok := cache.Get("0"); v != "" || ok {
			t.Errorf("expected empty cache")
		}
		if v, ok := cache.Get("1"); v != "" || ok {
			t.Errorf("expected empty cache")
		}

	})
}
