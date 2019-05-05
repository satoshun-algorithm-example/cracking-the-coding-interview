class StringBuilder:
    __slots__ = ['count', 'value']

    def __init__(self, capacity):
        self.count = 0
        self.value = ['' for _ in range(capacity)]

    def append(self, s):
        if isinstance(s, int):
            return self._append_int(s)
        if isinstance(s, str):
            return self._append_string(s)
        if s is None:
            return self._append_none()
        raise Exception("not implemented type " + str(type(s)))

    def _append_int(self, i):
        if i is None:
            return self._append_none()

        size = Int.string_size(i)
        if i < 0:
            size += 1
        self._ensureCapacityInternal(self.count + size)
        Int.get_chars(i, self.count + size, self.value)
        self.count += size

        return self

    def _append_string(self, s):
        if s is None:
            return self._append_none()
        l = len(s)
        self._ensureCapacityInternal(self.count + l)
        String.get_chars(
            s,
            0,
            l,
            self.value,
            self.count
        )
        self.count += l
        return self

    def _append_none(self):
        self._ensureCapacityInternal(self.count + 4)
        self.value[self.count] = 'n'
        self.value[self.count+1] = 'o'
        self.value[self.count+2] = 'n'
        self.value[self.count+3] = 'e'
        self.count += 4

    def _ensureCapacityInternal(self, size):
        pass

    def toString(self):
        print(self.value)
        if self.count == 0:
            return ''
        return StringFactory.newStringFromChars(0, self.count, self.value)

class StringFactory:
    @staticmethod
    def newStringFromChars(offset, count, value):
        s = ''
        for i in range(count):
            s += value[i]
        return s

class String:
    @staticmethod
    def get_chars(s, begin, end, dst, dst_begin):
        for j, i in enumerate(range(begin, end)):
            dst[dst_begin+j] = s[i]

class Int:
    size_table = [9, 99, 999, 9999, 99999]

    @staticmethod
    def string_size(x):
        if x < 0:
            x = -x
        for i in range(10000):
            if Int.size_table[i] >= x:
                return i + 1

    @staticmethod
    def get_chars(i, index, dst):
        sign = 0
        if i < 0:
            sign = '-'
            i = -i

        index -= 1
        while True:
            q = i // 10
            digit = i - q * 10
            dst[index] = str(digit)
            index -= 1
            i = q

            if q == 0:
                break

        if sign != 0:
            dst[index] = sign

builder = StringBuilder(32)

builder.append(None)
assert builder.toString() == 'none'

builder.append(' test ')
assert builder.toString() == 'none test '

builder.append(123)
assert builder.toString() == 'none test 123'

builder.append(-987)
assert builder.toString() == 'none test 123-987'
