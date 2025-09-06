package middleware

import (
	"bytes"
	"crypto/md5"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/tdewolff/minify/v2"
	"github.com/tdewolff/minify/v2/css"
)

const (
	TextCss           = "text/css"
	ContentType       = "Content-Type"
	CacheControl      = "Cache-Control"
	ETag              = "ETag"
	CacheControlValue = "public, max-age=31536000"
)

type CacheItem struct {
	Content   []byte
	ETag      string
	Timestamp time.Time
}

type Cache struct {
	items map[string]*CacheItem
	mutex sync.RWMutex
}

func NewCache() *Cache {
	return &Cache{
		items: make(map[string]*CacheItem),
	}
}

func (c *Cache) Get(key string) (*CacheItem, bool) {
	c.mutex.RLock()
	defer c.mutex.RUnlock()
	item, exists := c.items[key]
	return item, exists
}

func (c *Cache) Set(key string, item *CacheItem) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	c.items[key] = item
}

func (c *Cache) Clear() {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	c.items = make(map[string]*CacheItem)
}

var cache = NewCache()

func MinifyCSS() fiber.Handler {
	m := minify.New()
	m.AddFunc(TextCss, css.Minify)

	return func(c *fiber.Ctx) error {
		if !strings.HasPrefix(c.Path(), "/static/") {
			return c.Next()
		}

		ext := filepath.Ext(c.Path())
		if ext != ".css" {
			return c.Next()
		}

		filePath := "./public" + c.Path()
		fileInfo, err := os.Stat(filePath)
		if err != nil {
			return c.Next()
		}

		cacheKey := fmt.Sprintf("%s_%d", filePath, fileInfo.ModTime().Unix())
		if cachedItem, exists := cache.Get(cacheKey); exists {
			if c.Get("If-None-Match") == cachedItem.ETag {
				c.Status(304)
				return nil
			}

			c.Set(ContentType, TextCss)
			c.Set(CacheControl, CacheControlValue)
			c.Set(ETag, cachedItem.ETag)

			return c.Send(cachedItem.Content)
		}

		file, err := os.Open(filePath)
		if err != nil {
			return c.Next()
		}
		defer file.Close()

		content, err := io.ReadAll(file)
		if err != nil {
			return c.Next()
		}

		var minified bytes.Buffer
		if err := m.Minify(TextCss, &minified, bytes.NewReader(content)); err != nil {
			c.Set(ContentType, TextCss)
			return c.Send(content)
		}

		etag := generateMD5ETag(minified.Bytes())
		cacheItem := &CacheItem{
			Content:   minified.Bytes(),
			ETag:      etag,
			Timestamp: time.Now(),
		}
		cache.Set(cacheKey, cacheItem)

		c.Set(ContentType, TextCss)
		c.Set(CacheControl, CacheControlValue)
		c.Set(ETag, etag)

		return c.Send(minified.Bytes())
	}
}

func generateMD5ETag(content []byte) string {
	hash := md5.Sum(content)
	return fmt.Sprintf(`"%x"`, hash)
}
