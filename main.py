from json import loads
from urllib3 import disable_warnings
from os import environ
from logging import basicConfig, info
from sys import stdout
from distutils.util import strtobool
# Подключение самописных либ
from src.ImportCustomEnv import envCluster
from src.WorkWithApi import WorkWithApi
from src.ComparisonOfInformation import ComparisonOfInformation

envCluster()

RMQ_PORT = environ['RMQ_PORT']
RMQ_HOST = environ['RMQ_HOST']
RMQ_USER = environ['RMQ_USER']
RMQ_PASSWORD = environ['RMQ_PASSWORD']
RMQ_OBJECT = environ['RMQ_OBJECT'].split(', ')
RMQ_PATH_TO_DEFINITIONS_JSON_FILE = environ['RMQ_PATH_TO_DEFINITIONS_JSON_FILE']
RMQ_UPDATE_DEFINITIONS = environ['RMQ_UPDATE_DEFINITIONS']
RMQ_COMPARE_INFORMATION = environ['RMQ_COMPARE_INFORMATION']
RMQ_LOGGING_LEVEL = environ['RMQ_LOGGING_LEVEL']

if __name__ == "__main__":
    disable_warnings()
    definitions = 'definitions'

    # Настройки логгирования
    basicConfig(
        stream=stdout,
        level=RMQ_LOGGING_LEVEL,
        format='%(asctime)s.%(msecs)03d %(levelname)s  %(message)s',
        datefmt='%Y-%m-%d %H:%M:%S'
    )

    info("Connect to https://%s:%s" % (RMQ_HOST, RMQ_PORT))

    # Обновление дефинишинов
    if strtobool(RMQ_UPDATE_DEFINITIONS):
        if WorkWithApi(host=RMQ_HOST, port=RMQ_PORT, password=RMQ_PASSWORD, user=RMQ_USER,rmq_obj=definitions).GetRequest():
            info("GET INFORMATION FROM API ABOUT OBJECT: %s " % definitions)
            WorkWithApi(host=RMQ_HOST, port=RMQ_PORT, password=RMQ_PASSWORD, user=RMQ_USER,rmq_obj=definitions)\
                .UpdateDefinitions(file_definitions=RMQ_PATH_TO_DEFINITIONS_JSON_FILE)

    # Сравнение дефинишинов
    if strtobool(RMQ_COMPARE_INFORMATION):
        for obj in RMQ_OBJECT:
            if obj != definitions:
                if WorkWithApi(host=RMQ_HOST, port=RMQ_PORT, password=RMQ_PASSWORD, user=RMQ_USER, rmq_obj=obj).GetRequest():
                    info("GET INFORMATION FROM API ABOUT OBJECT: %s " % obj)

                    # Получение информации через api объектов полученных из env RMQ_OBJECT и перекладывание данных в json
                    data_definitions_from_api = loads(
                        WorkWithApi(host=RMQ_HOST, port=RMQ_PORT, password=RMQ_PASSWORD, user=RMQ_USER,rmq_obj=obj.lower()).DumpToJson())

                    with open(RMQ_PATH_TO_DEFINITIONS_JSON_FILE) as file:  # Открытие файла с новыми дефинишинами для создания переменной
                        new_data_definitions_file_local = loads(file.read())

                    # Запуск метода по сравнению дефинишинов
                    ComparisonOfInformation(rmq_obj=obj.lower()).DiffDefinitions(new_data_definitions_file_local=new_data_definitions_file_local,
                    data_definitions_from_api=data_definitions_from_api)
