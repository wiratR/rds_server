// ignore_for_file: use_build_context_synchronously, deprecated_member_use

import 'dart:convert';
import 'package:flutter/material.dart';
import 'package:flutter/services.dart';
import 'package:url_launcher/url_launcher.dart';
import '../../../../components/appbar/custom_app_bar.dart';
import '../../../../service/ksher/ksherpay_service.dart';
import '../../../../utils/utils.dart';

class TopUpScreen extends StatefulWidget {
  @override
  _TopUpScreenState createState() => _TopUpScreenState();
}

class _TopUpScreenState extends State<TopUpScreen> {
  int? selectedAmount;
  bool isCardSelected = false;
  bool isPromtPaySelected = false;
  bool isBankTransferSelected = false;
  String privateKeyContent = '';
  String publicKeyContent = '';
  bool isLoading = false;

  @override
  void initState() {
    super.initState();
    _loadAsset();
  }

  void _loadAsset() async {
    try {
      privateKeyContent =
          await rootBundle.loadString('assets/key/Mch38806_PrivateKey.pem');
      publicKeyContent = '''
-----BEGIN PUBLIC KEY-----
MFwwDQYJKoZIhvcNAQEBBQADSwAwSAJBAL7955OCuN4I8eYNL/mixZWIXIgCvIVE
ivlxqdpiHPcOLdQ2RPSx/pORpsUu/E9wz0mYS2PY7hNc2mBgBOQT+wUCAwEAAQ==
-----END PUBLIC KEY-----''';
      setState(() {}); // Trigger a rebuild to display the content
    } catch (e) {
      debugPrint('Error loading asset: $e');
    }
  }

  Future<void> _handleGateWayPay(
    String privateKey,
    String publicKey,
    bool isCardSelected,
    bool isPromtPaySelected,
    bool isBankTransferSelected,
    int selectedAmount,
    BuildContext context,
  ) async {
    Client client = Client(
      appId: "mch38806",
      privateKey: privateKey,
      publicKey: publicKey,
    );

    try {
      KsherPayContentResponse response = await client.gateWayPay(
        mchOrderNo: generateMchOrderNo(),
        feeType: 'THB',
        channelList: isCardSelected
            ? 'card'
            : isBankTransferSelected
                ? 'bank_transfer'
                : '',
        mchCode: generateMchOrderNo(),
        mchRedirectUrl: 'http://localhost/topup/payment-complete/',
        mchRedirectUrlFail: 'http://localhost/topup/payment-failed/',
        productName: 'Tapandpay',
        referUrl: 'http://localhost/topup/',
        device: 'sample',
        totalFee: selectedAmount,
      );
      debugPrint(response.toString());
      debugPrint(response.data.payContent);

      _launchURL(context, response.data.payContent);
    } catch (e) {
      debugPrint('Error: $e');
    } finally {
      setState(() {
        isLoading = false;
      });
    }
  }

  Future<void> _handleNativePay(
    String privateKey,
    String publicKey,
    bool isCardSelected,
    bool isPromtPaySelected,
    bool isBankTransferSelected,
    int selectedAmount,
    BuildContext context,
  ) async {
    Client client = Client(
      appId: "mch38806",
      privateKey: privateKey,
      publicKey: publicKey,
    );

    try {
      KsherResp response = await client.nativePay(
        mchOrderNo: generateMchOrderNo(),
        feeType: 'THB',
        channel: 'promptpay',
        totalFee: selectedAmount,
      );

      debugPrint(response.toString());

      if (isPromtPaySelected) {
        _showPromtPay(context, response);
      }
    } catch (e) {
      debugPrint('Error: $e');
    } finally {
      setState(() {
        isLoading = false;
      });
    }
  }

  void _launchURL(BuildContext context, String url) async {
    if (await canLaunch(url)) {
      await launch(url);
    } else {
      //
      ScaffoldMessenger.of(context).showSnackBar(
        SnackBar(content: Text('Could not launch $url')),
      );
    }
  }

  void _showPromtPay(BuildContext context, KsherResp responseMessage) {
    Map<String, String> stringMap = responseMessage.data
        .map((key, value) => MapEntry(key, value.toString()));
    String base64String = stringMap['imgdat']!.split(',').last;
    Uint8List bytes = base64Decode(base64String);

    showDialog(
      context: context,
      builder: (BuildContext context) {
        return AlertDialog(
          title: const Text("Payment Type 'ThaiQR'"),
          content: Column(
            mainAxisSize: MainAxisSize.min,
            children: [
              Image.memory(bytes),
              const SizedBox(height: 20),
              const Text(
                'please scan the Qr code using the mobilr app witn in 10 minutes.',
                style: TextStyle(fontSize: 14),
              ),
            ],
          ),
          actions: [
            TextButton(
              child: const Text('OK'),
              onPressed: () {
                Navigator.of(context).pop();
              },
            ),
          ],
        );
      },
    );
  }

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: CustomAppBar(
        logoAssetPath: 'assets/logo.png',
        showBackButton: true,
        showIconLogout: false,
      ),
      body: Padding(
        padding: const EdgeInsets.all(20.0),
        child: Column(
          crossAxisAlignment: CrossAxisAlignment.start,
          children: [
            const Text('Top-up Amount', style: TextStyle(fontSize: 18)),
            Wrap(
              spacing: 10,
              children: [100, 200, 500, 1000].map((amount) {
                return ChoiceChip(
                  label: Text('\$$amount'),
                  selected: selectedAmount == amount,
                  onSelected: (selected) {
                    setState(() {
                      selectedAmount = selected ? amount : null;
                    });
                  },
                );
              }).toList(),
            ),
            const SizedBox(height: 20),
            const Text('Select Payment Method', style: TextStyle(fontSize: 18)),
            CheckboxListTile(
              title: const Text('Credit/Debit Card'),
              value: isCardSelected,
              onChanged: (bool? value) {
                setState(() {
                  isCardSelected = value!;
                });
              },
            ),
            CheckboxListTile(
              title: const Text('Promtpay'),
              value: isPromtPaySelected,
              onChanged: (bool? value) {
                setState(() {
                  isPromtPaySelected = value!;
                });
              },
            ),
            CheckboxListTile(
              title: const Text('Bank Transfer'),
              value: isBankTransferSelected,
              onChanged: (bool? value) {
                setState(() {
                  isBankTransferSelected = value!;
                });
              },
            ),
            const SizedBox(height: 20),
            Center(
              child: ElevatedButton(
                onPressed: isLoading
                    ? null
                    : () async {
                        if (selectedAmount == null || selectedAmount! <= 0) {
                          debugPrint('Invalid amount selected.');
                          return;
                        }

                        if (!isCardSelected &&
                            !isPromtPaySelected &&
                            !isBankTransferSelected) {
                          debugPrint('No payment method selected.');
                          return;
                        }

                        setState(() {
                          isLoading = true;
                        });

                        if (isPromtPaySelected) {
                          try {
                            await _handleNativePay(
                              privateKeyContent,
                              publicKeyContent,
                              isCardSelected,
                              isPromtPaySelected,
                              isBankTransferSelected,
                              selectedAmount!,
                              context,
                            );
                          } catch (e) {
                            debugPrint('Error: $e');
                          }
                        } else {
                          try {
                            await _handleGateWayPay(
                              privateKeyContent,
                              publicKeyContent,
                              isCardSelected,
                              isPromtPaySelected,
                              isBankTransferSelected,
                              selectedAmount!,
                              context,
                            );
                          } catch (e) {
                            debugPrint('Error: $e');
                          }
                        }
                      },
                child: isLoading
                    ? const CircularProgressIndicator(
                        valueColor: AlwaysStoppedAnimation<Color>(Colors.white),
                      )
                    : const Text('Top Up'),
              ),
            ),
          ],
        ),
      ),
    );
  }
}
