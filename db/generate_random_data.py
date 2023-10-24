import random


if __name__ == '__main__':
    file = "data.sql"
    count = 1000
    filereader = open(file=file, mode='w')
    print(filereader.name)

    for i in range(0, count):
        input_a = random.random()
        input_b = random.random()
        output = "NULL"
        insert_statement = (("INSERT INTO `data` (`input_a`, `input_b`, `output`) " +
                            "VALUES ('{input_a}', '{input_b}', {output});\n")
                            .format(input_a=input_a, input_b=input_b, output=output))

        filereader.write(insert_statement)

