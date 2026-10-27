import subprocess
import sys
from pathlib import Path

SCRIPT = Path(__file__).resolve().parent.parent / "emailSlicer.py"


def run_slicer(stdin_text):
    result = subprocess.run(
        [sys.executable, str(SCRIPT)],
        input=stdin_text,
        capture_output=True,
        text=True,
        timeout=10,
    )
    assert result.returncode == 0
    return result.stdout.splitlines()


def test_valid_email():
    lines = run_slicer("john.doe@example.com\n")
    assert lines == [
        "Please enter your Email Id:",
        "Your username is:  john.doe",
        "Your domain is:  example.com",
    ]


def test_missing_at_sign():
    lines = run_slicer("not-an-email\n")
    assert lines == [
        "Please enter your Email Id:",
        "Please enter a valid Email Id.",
    ]


def test_surrounding_whitespace_stripped():
    lines = run_slicer("   jane@site.org \t\n")
    assert lines[1] == "Your username is:  jane"
    assert lines[2] == "Your domain is:  site.org"
