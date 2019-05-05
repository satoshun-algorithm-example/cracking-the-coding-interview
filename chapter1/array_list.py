class ArrayList:
    __slots__ = ['size', 'element_data']

    def __init__(self, capacity):
        self.size = 0
        self.element_data = [None for _ in range(capacity)]

    def add(self, e):
        self._ensure_capacity(self.size + 1)
        self.element_data[self.size] = e
        self.size += 1

    def get(self, index):
        return self.element_data[index]

    def _ensure_capacity(self, minimum_capacity):
        if minimum_capacity - len(self.element_data) > 0:
            self._grow(minimum_capacity)

    def _grow(self, minimum_capacity):
        new_size = minimum_capacity * 2
        new_array = [None for _ in range(new_size)]
        for i, e in enumerate(self.element_data):
            new_array[i] = e
        self.element_data = new_array

l = ArrayList(1)
l.add('test')
assert l.get(0) == 'test'

l.add('test2')
assert l.get(0) == 'test'
assert l.get(1) == 'test2'
