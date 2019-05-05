class StringBuilder:
    __slots__ = ['count', 'value']

    def __init__(self, capacity):
        self.count = 0
        self.value = ['' for _ in range(capacity)]

    def append(self, s):
        if isinstance(s, int):
            return
        if isinstance(s, str):
            return self._append_string(s)
        if s is None:
            return self._append_none()
        raise Exception("not implemented type " + str(type(s)))

    def _append_string(self, s):
        if s is None:
            return self._append_none()
        l = len(s)
        self._ensureCapacityInternal(self.count + l)
        String.getChars(
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
    def getChars(s, begin, end, dst, dst_begin):
        for j, i in enumerate(range(begin, end)):
            dst[dst_begin+j] = s[i]

builder = StringBuilder(16)

builder.append(None)
assert builder.toString() == 'none'

builder.append(' test')
assert builder.toString() == 'none test'
