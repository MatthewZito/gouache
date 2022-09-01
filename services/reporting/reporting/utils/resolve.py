"""Common utilities for resolving nested data"""


from typing import Dict


def resolve_page_key(result: Dict) -> str:
    """Extract the page key from a DynamoDB scan result if extant.

    Args:
        result (Dict): A DynamoDB scan result.

    Returns:
        str: The page key, or an empty string if not extant.
    """
    maybe_key = result.get('LastEvaluatedKey')
    if maybe_key is not None:
        return maybe_key.get('id')
    return ''
