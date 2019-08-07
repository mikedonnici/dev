def slices(series, length):
    if not len(series):
        raise ValueError("series arg should be a string")
    if length <= 0:
        raise ValueError("length arg should be greater than 0")
    if length > len(series):
        raise ValueError("length arg larger than series arg")

    s = [series[n:n + length] for n in range(0, len(series))]

    return list(filter(lambda x: len(x) >= length, s))
