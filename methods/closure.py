

def adder():
    sum = 0
    def func(x):
        nonlocal sum
        sum += x
        return sum
    return func

def main():
    print("a")
    pos, neg = adder(), adder()
    for i in range(10):
        print(pos(i), neg(-2*i))

if __name__ == '__main__':
    main()
