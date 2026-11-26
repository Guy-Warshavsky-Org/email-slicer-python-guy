"""Tests for :func:`emailSlicer.slice_email`."""

from emailSlicer import slice_email


def test_valid_email() -> None:
    assert slice_email("avimax37@gmail.com") == ("avimax37", "gmail.com")


def test_no_at_sign() -> None:
    assert slice_email("noatsign") is None


def test_multiple_at_signs_split_on_first() -> None:
    assert slice_email("a@b@c.com") == ("a", "b@c.com")
