import 'package:flutter/material.dart';

import '../../l10n/app_localizations.dart';
import '../controller/exchange_rate.dart';

class ExchangeRate extends StatefulWidget {
  final ExchangeRateController controller;
  final Function(String page) navigate;
  const ExchangeRate({super.key, required this.navigate, required this.controller});

  @override
  State<ExchangeRate> createState() => _ExchangeRateState();
}

class _ExchangeRateState extends State<ExchangeRate> {

  final TextEditingController _inputController = TextEditingController(text: '0.25');


  void sendExchangeRate(){

    String input = _inputController.text;

    double? value = double.tryParse(input);

    if (value != null) {
      widget.controller.addExchangeRate(value);
    }


  }


  @override
  Widget build(BuildContext context) {
    const primaryDarkColor = Color(0xFF131A2A);
    const borderColor = Color(0xFFE5E7EB);
    final l10n = AppLocalizations.of(context)!;


    return Center(
      child: ConstrainedBox(
        constraints: const BoxConstraints(
          maxWidth: 600,
        ),
        child: Padding(
          padding: const EdgeInsets.all(16.0),
          child: Column(
            crossAxisAlignment: CrossAxisAlignment.start,
            children: [
              Container(
                decoration: BoxDecoration(
                  color: const Color(0xFFF3F4F6),
                  borderRadius: BorderRadius.circular(8),
                ),
                child: TextButton.icon(
                  onPressed: () => widget.navigate("home"),
                  icon: const Icon(Icons.arrow_back, size: 18, color: primaryDarkColor),
                  label: Text(
                    l10n.backButton,
                    style: TextStyle(
                      color: primaryDarkColor,
                      fontWeight: FontWeight.w500,
                    ),
                  ),
                  style: TextButton.styleFrom(
                    padding: const EdgeInsets.symmetric(horizontal: 16, vertical: 12),
                    shape: RoundedRectangleBorder(
                      borderRadius: BorderRadius.circular(8),
                    ),
                  ),
                ),
              ),

              const SizedBox(height: 16),

              Container(
                padding: const EdgeInsets.all(24),
                decoration: BoxDecoration(
                  color: Colors.white,
                  border: Border.all(color: borderColor),
                  borderRadius: BorderRadius.circular(12),
                  boxShadow: [
                    BoxShadow(
                      color: Colors.black.withOpacity(0.02),
                      blurRadius: 8,
                      offset: const Offset(0, 4),
                    ),
                  ],
                ),
                child: Column(
                  crossAxisAlignment: CrossAxisAlignment.start,
                  children: [
                     Row(
                      children: [
                        Text(
                          '\$',
                          style: TextStyle(
                            fontSize: 24,
                            fontWeight: FontWeight.bold,
                            color: primaryDarkColor,
                          ),
                        ),
                        SizedBox(width: 8),
                        Text(
                           l10n.exchangeRateTitle,
                          style: TextStyle(
                            fontSize: 22,
                            fontWeight: FontWeight.bold,
                            color: primaryDarkColor,
                          ),
                        ),
                      ],
                    ),
                    const SizedBox(height: 8),

                    Text(
                      l10n.exchangeRateDesc,
                      style: TextStyle(
                        fontSize: 14,
                        color: Colors.grey.shade600,
                        height: 1.4,
                      ),
                    ),
                    const SizedBox(height: 24),

                    const Text(
                      '1 R\$ = USD',
                      style: TextStyle(
                        fontSize: 14,
                        fontWeight: FontWeight.bold,
                        color: primaryDarkColor,
                      ),
                    ),
                    const SizedBox(height: 8),

                    TextField(
                      controller: _inputController,
                      keyboardType: const TextInputType.numberWithOptions(decimal: true),
                      decoration: InputDecoration(
                        contentPadding: const EdgeInsets.symmetric(horizontal: 16, vertical: 14),
                        enabledBorder: OutlineInputBorder(
                          borderRadius: BorderRadius.circular(8),
                          borderSide: const BorderSide(color: borderColor),
                        ),
                        focusedBorder: OutlineInputBorder(
                          borderRadius: BorderRadius.circular(8),
                          borderSide: const BorderSide(color: primaryDarkColor, width: 1.5),
                        ),
                      ),
                    ),
                    const SizedBox(height: 24),

                    SizedBox(
                      width: double.infinity,
                      child: ElevatedButton(
                        onPressed: () => sendExchangeRate(),
                        style: ElevatedButton.styleFrom(
                          backgroundColor: primaryDarkColor,
                          padding: const EdgeInsets.symmetric(vertical: 16),
                          elevation: 0,
                          shape: RoundedRectangleBorder(
                            borderRadius: BorderRadius.circular(8),
                          ),
                        ),
                        child: Text(
                          l10n.exchangeRateSave,
                          style: TextStyle(
                            color: Colors.white,
                            fontSize: 16,
                            fontWeight: FontWeight.bold,
                          ),
                        ),
                      ),
                    ),
                  ],
                ),
              ),
            ],
          ),
        ),
      ),
    );
  }
}