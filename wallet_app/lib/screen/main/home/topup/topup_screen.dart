import 'package:flutter/material.dart';
import 'package:flutter/services.dart';
import '../../../../components/appbar/custom_app_bar.dart';
import '../../../../service/ksher/ksherpay_service.dart';
import '../../../../utils/utils.dart';
import 'promtpay/promptpay_screen.dart';

class TopUpScreen extends StatefulWidget {
  @override
  _TopUpScreenState createState() => _TopUpScreenState();
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
    // Call the nativePay method with dynamic totalFee
    KsherResp response = await client.nativePay(
      mchOrderNo: generateMchOrderNo(),
      feeType: 'THB',
      channel: isPromtPaySelected
          ? 'promptpay'
          : isCardSelected
              ? 'card'
              : isBankTransferSelected
                  ? 'bank_transfer'
                  : '',
      totalFee: selectedAmount,
    );
    // // Handle the response
    // print('Response: ${response.toJson()}');
    // Handle the response and navigate to the new screen
    // print('Response: ${response.toJson()}');

    if (isPromtPaySelected) {
      Navigator.push(
        context,
        MaterialPageRoute(
          builder: (context) => PromptpayScreen(
            responseMessage: response,
          ),
        ),
      );
    }
  } catch (e) {
    print('Error: $e');
  }
}

class _TopUpScreenState extends State<TopUpScreen> {
  int? selectedAmount;
  bool isCardSelected = false;
  bool isPromtPaySelected = false;
  bool isBankTransferSelected = false;
  String privateKeyContent = '';
  String publicKeyContent = '';

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
      // await rootBundle.loadString('assets/key/public_key.pem');
      setState(() {}); // Trigger a rebuild to display the content
    } catch (e) {
      print('Error loading asset: $e');
    }
  }

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: CustomAppBar(
        logoAssetPath: 'assets/logo.png', // Path to your logo asset
        showBackButton: true,
        showIconLogout: false,
      ),
      body: Padding(
        padding: const EdgeInsets.all(20.0),
        child: Column(
          crossAxisAlignment: CrossAxisAlignment.start,
          children: [
            Text('Top-up Amount', style: TextStyle(fontSize: 18)),
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
            SizedBox(height: 20),
            Text('Select Payment Method', style: TextStyle(fontSize: 18)),
            CheckboxListTile(
              title: Text('Credit/Debit Card'),
              value: isCardSelected,
              onChanged: (bool? value) {
                setState(() {
                  isCardSelected = value!;
                });
              },
            ),
            CheckboxListTile(
              title: Text('Promtpay'),
              value: isPromtPaySelected,
              onChanged: (bool? value) {
                setState(() {
                  isPromtPaySelected = value!;
                });
              },
            ),
            CheckboxListTile(
              title: Text('Bank Transfer'),
              value: isBankTransferSelected,
              onChanged: (bool? value) {
                setState(() {
                  isBankTransferSelected = value!;
                });
              },
            ),
            SizedBox(height: 20),
            Center(
              child: ElevatedButton(
                onPressed: () async {
                  // Validate selectedAmount
                  if (selectedAmount == null || selectedAmount! <= 0) {
                    print('Invalid amount selected.');
                    // Optionally, show a snackbar or dialog to inform the user
                    return;
                  }

                  // Validate at least one payment method selected
                  if (!isCardSelected &&
                      !isPromtPaySelected &&
                      !isBankTransferSelected) {
                    print('No payment method selected.');
                    // Optionally, show a snackbar or dialog to inform the user
                    return;
                  }
                  // Handle top-up logic here
                  print('Top-up Amount: $selectedAmount');
                  print('Credit/Debit Card: $isCardSelected');
                  print('Promtpay: $isPromtPaySelected');
                  print('Bank Transfer: $isBankTransferSelected');

                  try {
                    // Handle native payment logic with selected options
                    await _handleNativePay(
                      privateKeyContent,
                      publicKeyContent,
                      isCardSelected,
                      isPromtPaySelected,
                      isBankTransferSelected,
                      selectedAmount!,
                      context, // Pass context here
                    );
                  } catch (e) {
                    // Log error and provide user feedback
                    print('Error: $e');
                    // Optionally, show a snackbar or dialog to inform the user
                  }
                },
                child: Text('Top Up'),
              ),
            ),
          ],
        ),
      ),
    );
  }
}