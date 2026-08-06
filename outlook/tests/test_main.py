from unittest import TestCase
from unittest.mock import patch

from app import main


class StreamableHTTPServerTest(TestCase):
    def test_server_path_matches_oauth_proxy_target(self):
        with (
            patch.object(main, "setup_server") as setup_server_mock,
            patch.object(main.mcp, "run") as run,
        ):
            main.streamable_http_server()

        setup_server_mock.assert_called_once_with()
        run.assert_called_once_with(
            transport="streamable-http",
            host="0.0.0.0",
            port=9000,
            path="/mcp/outlook/",
        )
