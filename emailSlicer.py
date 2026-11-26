"""Email Slicer.

Prompts for an email address and prints its username and domain, splitting on
the first ``@`` character. Requires Python 3.10+.
"""


def slice_email(email: str) -> tuple[str, str] | None:
    """Split an email address into its username and domain.

    The split happens on the first ``@``; any later ``@`` stays in the domain.

    Args:
        email: The email address to slice.

    Returns:
        A ``(username, domain)`` tuple, or ``None`` if ``email`` has no ``@``.
    """
    # if email contains @ then valid email id otherwise invalid email id
    if email.find("@") != -1:
        username = email[:email.index("@")]     # username value = from 0 to index < @
        domain = email[email.index("@") + 1:]   # domain value = from index > @ to end
        return username, domain
    return None


def main() -> None:
    """Prompt for an email address and print its username and domain."""
    print("Please enter your Email Id:")
    email = input().strip()

    result = slice_email(email)
    if result is not None:
        username, domain = result
        print("Your username is: ", username)
        print("Your domain is: ", domain)
    else:
        print("Please enter a valid Email Id.")


if __name__ == "__main__":
    main()
