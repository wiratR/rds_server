class RequestPayment {
  final String appId;
  final int authCode;
  final String channel;
  final String feeType;
  final String mchOrderNo;
  final String nonceStr;
  final String timeStamp;
  final int totalFee;

  RequestPayment({
    required this.appId,
    required this.authCode,
    required this.channel,
    required this.feeType,
    required this.mchOrderNo,
    required this.nonceStr,
    required this.timeStamp,
    required this.totalFee,
  });

  factory RequestPayment.fromJson(Map<String, dynamic> json) {
    return RequestPayment(
      appId: json['appid'],
      authCode: json['auth_code'],
      channel: json['channel'],
      feeType: json['fee_type'],
      mchOrderNo: json['mch_order_no'],
      nonceStr: json['nonce_str'],
      timeStamp: json['time_stamp'],
      totalFee: json['total_fee'],
    );
  }

  Map<String, dynamic> toJson() {
    return {
      'appid': appId,
      'auth_code': authCode,
      'channel': channel,
      'fee_type': feeType,
      'mch_order_no': mchOrderNo,
      'nonce_str': nonceStr,
      'time_stamp': timeStamp,
      'total_fee': totalFee,
    };
  }
}
