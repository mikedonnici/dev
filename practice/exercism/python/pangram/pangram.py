def is_pangram(sentence):
    for r in range(ord('a'), ord('z') + 1):
        if not chr(r) in sentence.lower():
            return False

    return True
