from logging import info, error

'''
Класс предназначен для сравнения информации об объекте из API и файла с новыми дефинишенами
'''

# тк ключи для сравнения у некторых объектов разные приходится создавать массив с ключами и далее по ним сравнивать объекты из API и файла с новыми дефинишинами если эти ключи имеются
keysForCompare = ['name', 'source', 'permissions']


class ComparisonOfInformation(object):
    def __init__(self, rmq_obj=''):
        self.rmq_obj = rmq_obj

    def ReadFileWithNewDefinitions(self, new_data_definitions_file_local):
        # Создание массива со значениями из файла с новыми дефинишинами
        new_data = []
        for obj in new_data_definitions_file_local[self.rmq_obj]:
            for name in keysForCompare:
                if name in obj:
                    new_data.append(obj[name])
        return new_data

    def ReadDefinitionsFromApi(self, data_definitions_from_api):
        # Нахождение значения из API по ключам из keysForCompare
        new_data = []
        for obj in data_definitions_from_api:
            for name in keysForCompare:
                if name in obj:
                    new_data.append(obj[name])
        return new_data

    def DiffDefinitions(self, new_data_definitions_file_local, data_definitions_from_api):
        rmq_obj = self.rmq_obj

        # Нахождение разницы между файлом и хостом
        difference_values_from_local_file_definitions_and_host = set(
            self.ReadFileWithNewDefinitions(new_data_definitions_file_local=new_data_definitions_file_local)).difference(
            self.ReadDefinitionsFromApi(data_definitions_from_api=data_definitions_from_api))

        if difference_values_from_local_file_definitions_and_host:
            error("THE INFORMATION IS IN THE FILE BUT NOT IN THE HOST, FOR OBJECT %s: %s" % (
                rmq_obj, str(difference_values_from_local_file_definitions_and_host)))
            exit(1)
            return difference_values_from_local_file_definitions_and_host
        else:
            info("THE INFORMATION BETWEEN LOCAL FILE AND HOST IS EQUALS, FOR OBJECT: %s" % rmq_obj)

