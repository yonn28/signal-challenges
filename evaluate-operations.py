
def arithmeticExpression(a, b, c):
    return eval("a + b") == c or eval("a - b") == c or eval("a * b") == c or eval("a / b") == c


print(arithmeticExpression(5, 2, 5))