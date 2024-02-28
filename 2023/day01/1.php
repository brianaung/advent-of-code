<?php

$data = file('input.txt', FILE_SKIP_EMPTY_LINES);

$sum = 0;
foreach ($data as $line) {
    $nums = [];
    foreach (str_split($line) as $c) {
        if (ctype_digit($c)) {
            array_push($nums, $c);
        }
    }
	$n = $nums[0] . $nums[count($nums) - 1];
	$sum += intval($n);
}
echo $sum;
