import 'dart:convert';
import 'dart:math';
import 'package:date_format/date_format.dart';
import 'package:fast_rsa/fast_rsa.dart';
import 'package:flutter/foundation.dart';
import 'package:http/http.dart' as http;
import '../../utils/utils.dart';
import 'package:convert/convert.dart';

const String payDomain = "https://api.mch.ksher.net/KsherPay";
const String gateDomain = "https://gateway.ksher.com/api";
const String version = "v3.0.0"; // SDK version

void debugPrintResponseStatusCode(http.Response response) {
  debugPrint('response statusCode: ${response.statusCode}');
}

class HttpException implements Exception {
  final String message;
  HttpException(this.message);

  @override
  String toString() {
    return 'HttpException: $message';
  }
}

KsherResp convertDynamicToClass(dynamic data) {
  if (data is Map<String, dynamic>) {
    return KsherResp.fromJson(data);
  } else {
    throw ArgumentError('Expected a Map<String, dynamic>');
  }
}

class Client {
  final String appId;
  final String privateKey;
  final String publicKey;

  Client({
    required this.appId,
    required this.privateKey,
    required this.publicKey,
  });

  // Factory constructor to create a new Client instance
  factory Client.newClient(String appId, String privateKey, String publicKey) {
    debugPrint("--create new client ---\n");
    debugPrint("app id : $appId\n");
    debugPrint("privateKey : $privateKey\n");
    debugPrint("publicKey : $publicKey\n");

    return Client(
      appId: appId,
      privateKey: privateKey,
      publicKey: publicKey,
    );
  }

  Future<KsherResp> nativePay({
    required String mchOrderNo,
    required String feeType,
    required String channel,
    required int totalFee,
    String? redirectUrl,
    String? notifyUrl,
    String? paypageTitle,
    String? product,
    String? attach,
    String? operatorId,
    String? deviceId,
    String? imgType,
  }) {
    var postValue = {
      'appid': appId,
      'nonce_str': getNonceStr(4),
      'time_stamp': getTimeStamp(),
      'mch_order_no': mchOrderNo,
      'total_fee': totalFee.toString(),
      'fee_type': feeType,
      'channel': channel,
    };

    if (redirectUrl != null) postValue['redirect_url'] = redirectUrl;
    if (notifyUrl != null) postValue['notify_url'] = notifyUrl;
    if (paypageTitle != null) postValue['paypage_title'] = paypageTitle;
    if (product != null) postValue['product'] = product;
    if (attach != null) postValue['attach'] = attach;
    if (operatorId != null) postValue['operator_id'] = operatorId;
    if (deviceId != null) postValue['device_id'] = deviceId;
    if (imgType != null) postValue['img_type'] = imgType;

    return ksherPost("$payDomain/native_pay", postValue, privateKey, publicKey);
  }

  // https://gateway.ksher.com/api/gateway_pay
  Future<KsherPayContentResponse> gateWayPay({
    required String mchOrderNo,
    required String feeType,
    // alipay,wechat,linepay,airpay,promptpay,truemoney,card,ktccard,ktc_instal
    required String channelList,
    required String mchCode,
    // After the payment is succeed,
    // the web page is redirected to this URL.
    // If merchant leaves this field blank,
    // the web page will stay on payment page after payment succeed.
    required String mchRedirectUrl,
    required String mchRedirectUrlFail,
    required String productName,
    required String referUrl,
    required String device,
    required int totalFee,
  }) {
    var postValue = {
      'appid': appId,
      'nonce_str': getNonceStr(4),
      'time_stamp': getTimeStamp(),
      'mch_order_no': mchOrderNo,
      'fee_type': feeType,
      'channel_list': channelList,
      'mch_code': mchCode,
      'mch_redirect_url': mchRedirectUrl,
      'mch_redirect_url_fail': mchRedirectUrlFail,
      'product_name': productName,
      'refer_url': referUrl,
      'device': device,
      'total_fee': totalFee.toString()
    };

    // debugPrint(postValue.toString());
    return ksherPostGateWay(
        "$gateDomain/gateway_pay", postValue, privateKey, publicKey);
  }
}

class KsherResp {
  int code;
  String msg;
  String statusCode;
  String statusMsg;
  String sign; // Hexadecimal string
  String version;
  String timeStamp;
  Map<String, dynamic> data;

  KsherResp({
    required this.code,
    required this.msg,
    required this.statusCode,
    required this.statusMsg,
    required this.sign,
    required this.version,
    required this.timeStamp,
    required this.data,
  });

  factory KsherResp.fromJson(Map<String, dynamic> json) {
    return KsherResp(
      code: json['code'],
      msg: json['msg'],
      statusCode: json['status_code'],
      statusMsg: json['status_msg'],
      sign: json['sign'],
      version: json['version'],
      timeStamp: json['time_stamp'],
      data: json['data'] ?? {}, // Handle null case for 'data'
    );
  }

  Map<String, dynamic> toJson() {
    return {
      'code': code,
      'msg': msg,
      'status_code': statusCode,
      'status_msg': statusMsg,
      'sign': sign,
      'version': version,
      'time_stamp': timeStamp,
      'data': data,
    };
  }

  @override
  String toString() {
    return 'KsherResp{code: $code, msg: $msg, statusCode: $statusCode, statusMsg: $statusMsg, sign: $sign, version: $version, timeStamp: $timeStamp, data: $data}';
  }
}

class KsherPayContentResponse {
  final int code;
  final String msg;
  final PayContentData data;
  final String sign;
  final String message;

  KsherPayContentResponse({
    required this.code,
    required this.msg,
    required this.data,
    required this.sign,
    required this.message,
  });

  factory KsherPayContentResponse.fromJson(Map<String, dynamic> json) {
    return KsherPayContentResponse(
      code: json['code'],
      msg: json['msg'],
      data: PayContentData.fromJson(json['data']),
      sign: json['sign'],
      message: json['message'],
    );
  }

  Map<String, dynamic> toJson() {
    return {
      'code': code,
      'msg': msg,
      'data': data.toJson(),
      'sign': sign,
      'message': message,
    };
  }
}

class PayContentData {
  final String payContent;

  PayContentData({
    required this.payContent,
  });

  factory PayContentData.fromJson(Map<String, dynamic> json) {
    return PayContentData(
      payContent: json['pay_content'],
    );
  }

  Map<String, dynamic> toJson() {
    return {
      'pay_content': payContent,
    };
  }
}

String getNonceStr(int num) {
  const letters =
      'abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789';
  final random = Random();
  final buffer = StringBuffer();

  for (int i = 0; i < num; i++) {
    final index = random.nextInt(letters.length);
    buffer.write(letters[index]);
  }

  return buffer.toString();
}

String getTimeStamp() {
  final DateTime now = DateTime.now();
  return formatDate(now, [yyyy, mm, dd, HH, nn, ss]);
}

/*
签名
1.参数名排序
2.key1=valuekey2=value
3.appid=mch20027auth_code=12345channel=wechatfee_type=THBmch_order_no=2019051614001nonce_str=BpLnoperator_id=001time_stamp=20190517174933total_fee=100
*/
Future<String> ksherSign(
    Map<String, String> messageJson, String privateKeyPem) async {
  String strTosign = concatJson(messageJson);
  debugPrint("message for sign  = $strTosign\n");
  debugPrint("private Key = $privateKeyPem\n");
  var signature = await RSA.signPKCS1v15(strTosign, Hash.MD5, privateKeyPem);
  debugPrint("signature = $signature\n");
  // Convert the signature to bytes
  List<int> signatureBytes = base64.decode(signature);
  // Convert the bytes to a hexadecimal string using the custom function
  String hexSignature = bytesToHex(signatureBytes);
  debugPrint("sign (Hex) = $hexSignature\n");
  return hexSignature;
}

Future<bool> ksherVerify(message, String pubkey) async {
  KsherResp res = convertDynamicToClass(message);
  Map<String, String> stringMap =
      res.data.map((key, value) => MapEntry(key, value.toString()));
  String strToVerify = concatJson(stringMap);
  List<int> signatureBytes = hex.decode(res.sign);
  String decodedString = String.fromCharCodes(signatureBytes);

  debugPrint('sign : $decodedString');
  debugPrint('data : $strToVerify');
  debugPrint('public key : $pubkey');

  // debugdebugPrint(decodedString);
  // debugdebugPrint(strToVerify);

  try {
    bool isVerified =
        await RSA.verifyPKCS1v15(res.sign, strToVerify, Hash.MD5, pubkey);

    if (isVerified) {
      debugPrint('Sign verification Passed...');
      return true;
    } else {
      debugPrint('Sign verification failed...');
      return false;
    }
  } catch (e) {
    debugPrint('Sign verification failed...');
    return false;
  }
}

Future<KsherResp> ksherPost(
  String url,
  Map<String, String> postValue,
  String privateKeyData,
  String publicKeyData,
) async {
  KsherResp response = KsherResp(
    code: -1,
    msg: '',
    statusCode: '',
    statusMsg: '',
    sign: '',
    version: '',
    timeStamp: '',
    data: {},
  );
  try {
    // Sign the post data
    String sign = await ksherSign(postValue, privateKeyData);
    postValue['sign'] = sign;

    // Create the request
    var uri = Uri.parse(url);
    debugPrint("url for post : $uri\n");
    debugPrint("post value  : $postValue\n");
    var response = await http.post(
      uri,
      headers: {'Content-Type': 'application/x-www-form-urlencoded'},
      body: postValue,
    );

    debugPrint('Response : ${json.decode(response.body)}');

    if (response.statusCode == 200) {
      final responseData = json.decode(response.body);
      // verify sign
      if (responseData['code'] == 0) {
        // debugPrint('start --- verify sign -- ');
        // bool isValid = await ksherVerify(responseData, publicKeyData);
        // if (isValid) {
        //   return KsherResp.fromJson(responseData);
        // } else {
        //   throw Exception('Signature verification failed');
        // }
        return KsherResp.fromJson(responseData);
      }
    } else {
      debugPrint('Error: ${response.statusCode} - ${response.reasonPhrase}');
      throw Exception('Failed to post data');
    }
  } catch (e) {
    debugPrint('Error: $e');
  }

  return response;
}

Future<KsherPayContentResponse> ksherPostGateWay(
  String url,
  Map<String, String> postValue,
  String privateKeyData,
  String publicKeyData,
) async {
  PayContentData payContentData = PayContentData(payContent: '');
  KsherPayContentResponse response = KsherPayContentResponse(
    code: -1,
    msg: '',
    data: payContentData,
    sign: '',
    message: '',
  );
  try {
    // Sign the post data
    String sign = await ksherSign(postValue, privateKeyData);
    postValue['sign'] = sign;

    // Create the request
    var uri = Uri.parse(url);
    debugPrint("url for post : $uri\n");
    debugPrint("post value  : $postValue\n");
    var response = await http.post(
      uri,
      headers: {'Content-Type': 'application/x-www-form-urlencoded'},
      body: postValue,
    );

    debugPrint('Response : ${json.decode(response.body)}');

    if (response.statusCode == 200) {
      final responseData = json.decode(response.body);
      // verify sign
      if (responseData['code'] == 0) {
        // debugPrint('start --- verify sign -- ');
        // bool isValid = await ksherVerify(responseData, publicKeyData);
        // if (isValid) {
        //   return KsherPayContentResponse.fromJson(responseData);
        // } else {
        //   throw Exception('Signature verification failed');
        // }
        return KsherPayContentResponse.fromJson(responseData);
      }
    } else {
      debugPrint('Error: ${response.statusCode} - ${response.reasonPhrase}');
      throw Exception('Failed to post data');
    }
  } catch (e) {
    debugPrint('Error: $e');
  }

  return response;
}
