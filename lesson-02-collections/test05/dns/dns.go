package dns

import ("errors"
        "strings")

// ErrInvalid 表示域名或 IP 为空（去掉空白后）。
var ErrInvalid = errors.New("dns: domain and ip must not be empty")

// Resolver 是一份本地 DNS 缓存：域名 -> IP，同时支持 IP -> 域名的反查。
//
// 字段怎么设计由你决定，这一题是从零实现。
// 要求零值 Resolver（var r Resolver）也能直接使用。
type Resolver struct {
	forward map[string]string
	reverse map{string][]string
	// TODO: 设计你需要的字段。
}

// New 返回一个空的 Resolver。
func New() *Resolver {
	// TODO
	return &Resolver{}
}

// Normalize 把域名规范化：去掉两端空白、转成小写、去掉末尾的一个 "."。
//
//	" WWW.Example.COM. " -> "www.example.com"
func Normalize(domain string) string {
	// TODO
	d := strings.TrimSpace(domain)
	d = strings.ToLower(d)
	d = strings.TrimSuffix(d, ".")
	return d
}

// Add 记录 domain -> ip。域名先规范化，IP 去掉两端空白。
// 任意一个为空时返回 ErrInvalid，并且不修改 Resolver。
// 同一个域名再次 Add 时覆盖旧 IP，反查结果也要跟着更新。
func (r *Resolver) Add(domain, ip string) error {
	// TODO
	d := Normalize(domain)
	i := strings.TrimSpace(ip)
	if d == "" || i == "" {
		return ErrInvalid
	}
	if r.forward == nil {
		r.forward = make(map[string]string)
	}
	if r.reverse == nil {
		r.reverse = make(map[string][]string)
	}
	if old, ok := r.forward[d]; ok && old != i {
		list := r.reverse[old]
		for idx, v := range list {
			if v == d {
				r.reverse[old] = append(list[:idx], list[idx+1:]...)
				break
			}
		}
	}
	r.forward[d] = i
	for _, v := range r.reverse[i] {
		if v == d {
			return nil
		}
	}
	r.reverse[i] = append(r.reverse[i], d)
	return nil
}

// Lookup 查询域名对应的 IP；查询前同样要规范化。不存在时返回 "", false。
func (r *Resolver) Lookup(domain string) (string, bool) {
	// TODO
	d := Normalize(domain)
	ip, ok := r.forward[d]
	if !ok {
		return "", false
	}
	return ip, true
}

// Remove 删除一个域名；不存在时什么也不做。反查结果也要同步删除。
func (r *Resolver) Remove(domain string) {
	d := Normalize(domain)

	old, ok := r.forward[d]
	if !ok {
		return
	}

	delete(r.forward, d)

	list := r.reverse[old]
	for idx, v := range list {
		if v == d {
			r.reverse[old] = append(list[:idx], list[idx+1:]...)
			break
		}
	}
	if len(r.reverse[old]) == 0 {
		delete(r.reverse, old)
	}
	// TODO
}

// Len 返回记录的域名数量。
func (r *Resolver) Len() int {
	// TODO
      return len(r.forward)
}

// ReverseLookup 返回所有指向 ip 的域名，按字典序升序。
// 没有时返回长度为 0 的结果。调用方修改返回的切片不能影响 Resolver。
func (r *Resolver) ReverseLookup(ip string) []string {
	// TODO
	return nil
}

// Domains 返回所有域名，按字典序升序。
func (r *Resolver) Domains() []string {
	// TODO
	return nil
}
