from os import environ


def envCluster():
    if not "RMQ_GET_ENV_FROM_OS" in environ:
        from dotenv import load_dotenv
        load_dotenv()
