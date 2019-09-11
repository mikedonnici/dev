class Matrix():
    def __init__(self, matrix_string):
        self.data = self.__matrix(matrix_string)

    def row(self, index):
        return self.data[index-1]

    def column(self, index):
        col = []
        for row in self.data:
            col.append(row[index-1])
        return col

    def __matrix(self, str):
        m = []
        rows = str.split("\n")
        for i in range(len(rows)):
            m.append([int(n) for n in rows[i].split(" ")])
        return m
