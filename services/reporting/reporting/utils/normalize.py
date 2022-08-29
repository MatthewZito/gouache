def normalize_dynamo_report(result_item: object):
    return {
        'id': result_item.get('Id'),
        'data': result_item.get('Data'),
        'caller': result_item.get('Caller'),
        'name': result_item.get('Name'),
        'ts': result_item.get('TS'),
    }
