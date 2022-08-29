import unittest
import datetime
from flask import Flask, request
from analytics.repositories.base_repository import BaseRepository
from analytics.entities.report import Report
from analytics.repositories.__mocks__.report_repository_mock import MockReportRepository
from analytics.repositories.__mocks__.session_repository_mock import (
    MockSessionRepository,
)
from analytics.context.__mocks__.context_mock import test_context_set


class TestReportingController(unittest.TestCase):
    def setUp(self):
        self.app = Flask(__name__)

    def tearDown(self):
        pass

    def test_it(self):

        with test_context_set(app, MockReportRepository(), MockSessionRepository()):
            with self.app.test_client() as c:
                pass
                # resp = c.get('/users/me')
                # data = json.loads(resp.data)
                # self.assert_equal(data['username'], my_user.username)


if __name__ == '__main__':
    unittest.main()
