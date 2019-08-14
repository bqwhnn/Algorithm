package algorithm

func quickSort(nums []int) {
	len := len(nums)
	sort1(nums, 0, len-1)
}

func quickSort2(nums []int) {
	len := len(nums)
	sort2(nums, 0, len-1)
}

// 交换法
func sort1(a []int, first, last int) {
	if first < last {
		i, j, x := first+1, last, a[first]

		for i < j {
			for i < j && a[j] >= x {
				j--
			}
			for i < j && a[i] < x {
				i++
			}
			if i < j {
				a[i], a[j] = a[j], a[i]
			}
		}
		if a[first] > a[j] {
			a[first], a[j] = a[j], a[first]
		} else {
			j--
		}
		sort1(a, first, j-1)
		sort1(a, j+1, last)
	}
}

// 挖坑法
func sort2(a []int, first, last int) {
	if first < last {
		i, j, x := first, last, a[first]

		for i < j {
			for i < j && a[j] >= x {
				j--
			}
			if i < j {
				a[i] = a[j]
				i++
			}
			for i < j && a[i] < x {
				i++
			}
			if i < j {
				a[j] = a[i]
				j--
			}
		}
		a[i] = x
		sort2(a, first, i-1)
		sort2(a, i+1, last)
	}
}

/*
// Jon Bently beautiful code
void quicksort(int l, int u){
    int i, m;
    if(l >= u) return;
    m = l;
    for(i = l+1; i<= u; i++)
        if(x[i] < x[l]) // buggy!
            swap(++m, i);

    swap(l, m);

    quicksort(l, m-1);
    quicksort(m+1, u);

}
*/
