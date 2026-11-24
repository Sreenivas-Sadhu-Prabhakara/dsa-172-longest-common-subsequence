# Longest Common Subsequence

> Striver SDE Sheet 172 · Day 25 DP — the foundational string DP: match and step both back, or drop a character from one side; two rolling rows replace the table. Go, stdlib only, tested.

## Problem

The length of the longest sequence appearing (not necessarily contiguously, in order) in both strings — the backbone of diff tools, and the parent problem of edit distance, shortest common supersequence and half the string-DP family.

## Approach

State: the LCS of prefixes a[:i], b[:j]. If the last characters match, they belong to the answer — 1 + the LCS of both shortened. If not, the answer drops the last character of one side or the other — the max of the two. Since row i depends only on row i−1, two rolling rows (swapped by reference each iteration, never copied) shrink the table from O(nm) to O(m). Length-only is what the rolling rows buy; recovering the actual subsequence needs the full table to walk back through. The symmetry test (`LCS(a,b) == LCS(b,a)`) is a free property check.

**Complexity:** O(n·m) time, O(min(n,m)) space.

## Run

    go test ./...   # verify
    go run .        # demo

---

One of 191 repos solving Striver's SDE Sheet in Go — stdlib only, one problem per repo, every solution tested. Browse the set: [`dsa-*` repos](https://github.com/Sreenivas-Sadhu-Prabhakara?tab=repositories&q=dsa-&type=source).
