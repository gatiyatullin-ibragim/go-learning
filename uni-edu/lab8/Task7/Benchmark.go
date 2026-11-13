func TestCountParEqualsSeq(t *testing.T) {
    want := CountSeq(N)
    for _, parts := range []int{1, 2, 3, 7, 8, 1000, 200_000} {
        if got := CountPar(N, parts); got != want {
            t.Errorf("parts=%d: got %d, want %d", parts, got, want)
        }
    }
}