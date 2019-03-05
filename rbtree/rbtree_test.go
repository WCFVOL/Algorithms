package rbtree

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func less(k1, k2 interface{}) bool {
	return k1.(int) < k2.(int)
}

func equal(k1, k2 interface{}) bool {
	return k1 == k2
}

func TestRbtree_Insert(t *testing.T) {
	rand.Seed(time.Now().Unix())
	tree := NewRBTree(less, equal)
	list := [10000]int{}
	for i := 0; i < 10000; i++ {
		key := rand.Intn(10000)
		value := key
		tree.Insert(key, value)
		list[i] = key
	}
	for i := 9999; i >= 0; i-- {
		key := list[i]
		value := tree.Search(key)
		if key != value {
			fmt.Println(key, value)
		}
	}
}

func TestRbtree_Remove(t *testing.T) {
	//rand.Seed(time.Now().Unix())
	//tree := NewRBTree(less, equal)
	//for i := 0; i < 5; i++ {
	//	key := i
	//	value := key
	//	tree.Insert(key, value)
	//}
	//for i := 0; i < 5; i++ {
	//	key := i
	//	fmt.Println(tree.Remove(key))
	//}
	rand.Seed(time.Now().Unix())
	tree := NewRBTree(less, equal)
	list := [1000]int{601, 952, 515, 911, 808, 196, 632, 91, 573, 943, 148, 373, 459, 756, 401, 526, 468, 147, 494, 617, 326, 199, 867, 389, 502, 101, 121, 482, 225, 723, 315, 673, 759, 484, 821, 749, 713, 840, 773, 351, 249, 87, 462, 939, 77, 531, 936, 996, 402, 141, 376, 130, 679, 212, 952, 653, 361, 29, 32, 156, 390, 424, 163, 872, 172, 770, 760, 108, 92, 157, 246, 433, 969, 46, 942, 825, 429, 846, 740, 568, 460, 470, 374, 114, 768, 113, 177, 10, 309, 140, 620, 549, 872, 388, 700, 24, 954, 766, 407, 477, 768, 38, 420, 716, 74, 809, 859, 90, 485, 34, 388, 31, 174, 19, 789, 27, 373, 555, 697, 709, 754, 417, 459, 552, 1, 307, 150, 793, 811, 443, 553, 38, 999, 758, 621, 175, 374, 70, 853, 111, 182, 459, 452, 801, 691, 173, 569, 977, 136, 282, 804, 513, 637, 954, 854, 207, 781, 259, 235, 30, 445, 603, 114, 969, 874, 240, 880, 883, 548, 938, 674, 81, 772, 748, 575, 581, 394, 593, 862, 935, 905, 957, 511, 793, 700, 772, 566, 452, 715, 547, 26, 68, 304, 554, 750, 282, 864, 168, 885, 80, 321, 735, 530, 943, 954, 111, 402, 253, 218, 390, 367, 609, 671, 596, 701, 270, 622, 324, 164, 89, 963, 79, 518, 208, 309, 418, 822, 818, 402, 229, 35, 279, 969, 941, 889, 642, 38, 798, 567, 888, 566, 572, 580, 789, 529, 578, 688, 882, 767, 513, 254, 643, 741, 192, 16, 14, 629, 770, 599, 907, 967, 489, 911, 891, 141, 250, 921, 313, 27, 354, 26, 311, 488, 102, 167, 147, 190, 92, 375, 881, 907, 456, 282, 709, 480, 198, 468, 87, 647, 819, 122, 201, 499, 934, 980, 421, 491, 633, 312, 912, 685, 235, 900, 755, 579, 963, 665, 399, 504, 30, 541, 132, 604, 373, 184, 790, 127, 32, 400, 77, 736, 658, 715, 168, 476, 995, 604, 80, 819, 729, 991, 836, 758, 19, 61, 238, 254, 935, 872, 575, 733, 643, 510, 522, 919, 627, 616, 473, 652, 179, 853, 354, 796, 185, 716, 362, 563, 309, 873, 493, 182, 303, 141, 601, 272, 957, 19, 959, 783, 147, 472, 658, 976, 230, 679, 529, 13, 877, 859, 103, 228, 917, 318, 109, 313, 294, 713, 306, 338, 755, 787, 185, 260, 798, 43, 647, 885, 767, 215, 321, 557, 814, 364, 793, 157, 548, 796, 610, 161, 42, 325, 636, 899, 743, 179, 515, 196, 989, 265, 833, 366, 188, 1, 246, 126, 938, 966, 154, 419, 596, 219, 352, 500, 559, 608, 162, 408, 801, 917, 291, 627, 721, 371, 307, 992, 691, 895, 740, 935, 719, 18, 196, 24, 643, 627, 9, 532, 5, 641, 724, 930, 309, 454, 173, 381, 891, 329, 730, 782, 803, 227, 515, 501, 17, 292, 494, 522, 221, 624, 176, 602, 600, 823, 598, 540, 360, 100, 436, 275, 475, 257, 545, 306, 192, 463, 139, 845, 670, 156, 418, 480, 701, 580, 57, 54, 689, 886, 13, 349, 301, 170, 256, 329, 895, 68, 115, 287, 457, 880, 721, 946, 795, 574, 581, 558, 545, 268, 828, 798, 838, 448, 702, 910, 408, 218, 785, 232, 841, 888, 842, 49, 298, 208, 520, 905, 64, 941, 837, 527, 767, 496, 73, 249, 832, 896, 170, 299, 948, 191, 413, 798, 899, 719, 346, 511, 570, 187, 159, 234, 675, 810, 518, 55, 108, 551, 677, 379, 928, 270, 591, 917, 237, 95, 194, 739, 826, 671, 618, 747, 805, 523, 535, 613, 823, 997, 114, 143, 232, 825, 342, 588, 659, 463, 4, 699, 662, 864, 662, 543, 122, 847, 32, 771, 718, 87, 84, 817, 68, 353, 76, 230, 53, 705, 674, 854, 290, 802, 42, 913, 783, 50, 728, 946, 303, 138, 719, 325, 940, 983, 130, 795, 604, 896, 538, 498, 749, 581, 929, 617, 476, 817, 936, 987, 516, 265, 721, 250, 42, 777, 745, 519, 643, 760, 500, 214, 644, 955, 785, 424, 400, 992, 388, 443, 669, 922, 809, 302, 404, 334, 765, 8, 103, 868, 106, 482, 991, 312, 992, 337, 923, 103, 385, 298, 543, 199, 311, 738, 98, 840, 191, 542, 471, 972, 526, 118, 907, 37, 728, 553, 829, 518, 991, 453, 838, 811, 856, 694, 381, 722, 69, 759, 76, 99, 743, 104, 74, 705, 381, 779, 344, 910, 994, 32, 432, 102, 266, 968, 934, 282, 328, 840, 404, 331, 954, 585, 354, 481, 826, 305, 326, 23, 316, 702, 524, 929, 960, 233, 697, 874, 90, 391, 481, 682, 390, 805, 542, 575, 908, 743, 960, 639, 30, 11, 360, 936, 480, 347, 729, 771, 10, 113, 925, 882, 917, 850, 191, 830, 974, 73, 150, 603, 581, 493, 865, 599, 289, 480, 937, 965, 924, 383, 549, 473, 313, 871, 147, 840, 540, 521, 114, 137, 195, 954, 804, 903, 427, 910, 273, 25, 612, 786, 581, 38, 471, 749, 414, 922, 212, 249, 169, 157, 500, 183, 620, 718, 121, 576, 899, 850, 508, 100, 475, 670, 200, 316, 975, 349, 117, 249, 959, 70, 152, 170, 605, 471, 215, 60, 994, 470, 412, 811, 764, 171, 153, 657, 212, 83, 965, 956, 407, 123, 318, 593, 510, 972, 843, 359, 78, 326, 705, 765, 710, 621, 37, 477, 146, 952, 994, 892, 99, 129, 833, 627, 50, 140, 704, 493, 328, 844, 653, 329, 894, 556, 541, 361, 610, 567, 824, 955, 577, 695, 563, 94, 461, 356, 753, 569, 391, 882, 537, 654, 153, 267, 296, 321, 388, 946, 353, 731, 965, 563, 605, 740, 475, 75, 579, 582, 558, 480, 713, 804, 486, 639, 246, 323, 668, 597, 8, 805, 165, 278, 254, 393, 625, 929, 520, 285, 336, 698, 108, 529, 405, 638, 19, 46, 172, 388, 381, 967, 880, 699, 113, 149, 913, 431, 755, 673, 943, 566, 13, 958, 232, 710, 884, 868, 312, 941, 392, 125, 12, 490, 388, 364, 825, 383, 789, 850, 891, 259, 965}
	for i := 0; i < 1000; i++ {
		//key := rand.Intn(1000)
		key := list[i]
		value := key
		tree.Insert(key, value)
		list[i] = key
		fmt.Print(key, ",")
	}
	for i := 999; i >= 0; i-- {
		key := list[i]
		fmt.Println(key)
		value := tree.Remove(key)
		if key != value {
			fmt.Println(key, value)
		}
	}

}
