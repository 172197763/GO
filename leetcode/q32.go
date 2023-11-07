package leetcode

type s32 struct {
	bit     uint32
	str_len int
}

// https://leetcode.cn/problems/maximum-product-of-word-lengths/description/?envType=daily-question&envId=2023-11-06
func MaxProduct(words []string) int {
	res := 0
	bit_arr := make([]s32, 0)
	for _, v := range words {
		tmp := s32{0, len(v)}
		for i := 0; i < len(v); i++ {
			tmp.bit = tmp.bit | (1 << uint32(v[i]-97))
		}
		bit_arr = append(bit_arr, tmp)
	}
	for i := 0; i < len(bit_arr)-1; i++ {
		for j := i + 1; j < len(bit_arr); j++ {
			if bit_arr[i].bit&bit_arr[j].bit == 0 && res < bit_arr[i].str_len*bit_arr[j].str_len {
				res = bit_arr[i].str_len * bit_arr[j].str_len
			}
		}
	}
	return res
}
