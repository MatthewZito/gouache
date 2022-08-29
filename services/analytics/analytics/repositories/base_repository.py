from abc import ABC, abstractmethod


class BaseRepository(ABC):
    @abstractmethod
    def get(self, key: str):
        pass

    @abstractmethod
    def put(self, **kwargs):
        pass
