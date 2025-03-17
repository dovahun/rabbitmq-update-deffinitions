from requests import get, post, ConnectionError, ConnectTimeout, HTTPError
from json import dumps
from logging import info, error
from sys import exit

'''
Класс предназначен для работы с API
'''


class WorkWithApi(object):
    def __init__(self, host='', port='', user='', password='', rmq_obj=''):
        self.host = host
        self.port = port
        self.user = user
        self.password = password
        self.rmq_obj = rmq_obj

    def uri(self):
        url = 'https://%s:%s/api/%s' % (self.host, self.port, self.rmq_obj)
        return url

    def GetRequest(self):
        try:
            response = get(self.uri(), auth=(self.user, self.password), verify=False)
            if response.ok:
                return response
            else:
                error("NOT CORRECT KEY: %s" % self.rmq_obj)
                exit(1)
        except ConnectionError or ConnectTimeout:
            error("CONNECTION FAILED TO HOST: %s" % self.uri())
            exit(1)

    def DumpToJson(self):  # Получение информациии об объекте полученного через ENV RMQ_OBJECT через get запрос
        objects = self.GetRequest().json()
        json_dump_obj = dumps(objects)
        return json_dump_obj

    def UpdateDefinitions(self, file_definitions):  # Обновление дефинишинов через POST запрос
        headers = {'content-type': 'application/json'}
        with open(file_definitions) as file_definitions:
            response = post(
                url=self.uri(),
                auth=(self.user, self.password),
                verify=False,
                data=file_definitions,
                headers=headers
            )
        try:
            if response.ok:
                info("DEFINITIONS UPDATED")
                return response
        except HTTPError:
            error("UPDATE DEFINITIONS FAILED, EXIT %s" % response.status_code)
            exit(1)
