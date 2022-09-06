import requests

API_URL = "http://localhost:8086"

#def test_get_item_bad_request():
def test_get_item_not_found():
    pCode = "12345645"
    url = f"{API_URL}/item/{pCode}"
    response = requests.request("GET", url)
    assert 404 == response.status_code