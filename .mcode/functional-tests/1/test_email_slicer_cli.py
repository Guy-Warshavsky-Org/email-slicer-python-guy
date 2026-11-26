import subprocess, sys, pathlib, pytest

REPO = pathlib.Path(__file__).resolve().parents[3]
SCRIPT = REPO / "emailSlicer.py"
BASE = pathlib.Path(__import__("os").environ.get("MCODE_DIR", "/l2l/mcode")) / "baseline" / "before.txt"
PROMPT = "Please enter your Email Id:\n"


def run(stdin):
    return subprocess.run([sys.executable, str(SCRIPT)], input=stdin, capture_output=True, text=True, timeout=30)


def test_baseline_byte_identical():
    base = BASE.read_text()
    out = "".join(run(i + "\n").stdout for i in ["avimax37@gmail.com", "noatsign", "a@b@c.com"])
    assert out == base


def test_happy_path():
    r = run("avimax37@gmail.com\n")
    assert r.returncode == 0
    assert r.stdout == PROMPT + "Your username is:  avimax37\nYour domain is:  gmail.com\n"


def test_no_at():
    r = run("noatsign\n")
    assert r.returncode == 0 and r.stdout == PROMPT + "Please enter a valid Email Id.\n"


def test_multi_at():
    assert run("a@b@c.com\n").stdout == PROMPT + "Your username is:  a\nYour domain is:  b@c.com\n"


def test_whitespace_padded():
    assert run("   user@example.com \t \n").stdout == PROMPT + "Your username is:  user\nYour domain is:  example.com\n"


def test_empty_line():
    r = run("\n")
    assert r.returncode == 0 and r.stdout == PROMPT + "Please enter a valid Email Id.\n"


def test_whitespace_only():
    assert run("   \n").stdout == PROMPT + "Please enter a valid Email Id.\n"


def test_leading_at():
    assert run("@domain.com\n").stdout == PROMPT + "Your username is:  \nYour domain is:  domain.com\n"


def test_trailing_at():
    assert run("user@\n").stdout == PROMPT + "Your username is:  user\nYour domain is:  \n"


def test_only_at():
    assert run("@\n").stdout == PROMPT + "Your username is:  \nYour domain is:  \n"


def test_import_has_no_side_effects():
    r = subprocess.run([sys.executable, "-c", "import sys; sys.path.insert(0, %r); import emailSlicer as m; print(m.slice_email('a@b@c'), m.slice_email('x'))" % str(REPO)],
                       input="", capture_output=True, text=True, timeout=30)
    assert r.returncode == 0
    assert r.stdout == "('a', 'b@c') None\n"


def test_eof_no_stdin_behaves_like_original():
    r = run("")
    assert r.returncode != 0 and "EOFError" in r.stderr
