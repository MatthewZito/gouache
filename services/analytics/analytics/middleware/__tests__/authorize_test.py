from datetime import datetime
import unittest
from unittest import mock

from analytics.main import app
from analytics.repositories.base_repository import BaseRepository
from analytics.repositories.session_repository import SessionRepository


class MockReportRepository(BaseRepository):
    def __init__(self, table_name: str) -> None:
        self.tr = ""

    def get(self, key: str):
        pass

    def put(self, **kwargs):
        pass


class TestAuthorizationMiddleware(unittest.TestCase):
    def setUp(self):
        self.app = app

    def tearDown(self):
        pass

    @mock.patch(
        'analytics.context.context.SessionRepository',
    )
    @mock.patch('analytics.context.context.ReportRepository', new=MockReportRepository)
    def test_session_expired(self, m: mock.Mock):

        m.return_value = SessionRepository()
        # Avoid calling the constructor logic
        m.return_value.__init__ = mock.MagicMock()
        m.return_value.get = mock.MagicMock(
            return_value=f'{{"Username":"user","Expiry":"1022-08-29 21:59:59.999999"}}'
        )

        with self.app.app_context():
            with self.app.test_client() as c:
                c.set_cookie('gouache_session', 'gouache_session', '123')
                res = c.get(f'/api/report/c22c1173-93be-4550-9200-afe7df28bf2f')

                self.assertEqual(res.status_code, 401)

    @mock.patch(
        'analytics.context.context.SessionRepository',
    )
    @mock.patch('analytics.context.context.ReportRepository', new=MockReportRepository)
    def test_session_not_extant(self, m: mock.Mock):

        m.return_value = SessionRepository()
        # Avoid calling the constructor logic
        m.return_value.__init__ = mock.MagicMock()
        m.return_value.get = mock.MagicMock(return_value=None)

        with self.app.app_context():
            with self.app.test_client() as c:
                c.set_cookie('gouache_session', 'gouache_session', '123')
                res = c.get(f'/api/report/c22c1173-93be-4550-9200-afe7df28bf2f')

                self.assertEqual(res.status_code, 401)

    @mock.patch(
        'analytics.context.context.SessionRepository',
    )
    @mock.patch('analytics.context.context.ReportRepository', new=MockReportRepository)
    def test_cookie_not_extant(self, m: mock.Mock):

        m.return_value = SessionRepository()
        # Avoid calling the constructor logic
        m.return_value.__init__ = mock.MagicMock()
        m.return_value.get = mock.MagicMock(
            return_value=f'{{"Username":"user","Expiry":"{datetime.max}"}}'
        )

        with self.app.app_context():
            with self.app.test_client() as c:
                res = c.get(f'/api/report/c22c1173-93be-4550-9200-afe7df28bf2f')

                self.assertEqual(res.status_code, 401)


if __name__ == '__main__':
    unittest.main()
