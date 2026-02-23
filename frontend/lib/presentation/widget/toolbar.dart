import 'package:flutter/cupertino.dart';
import 'package:flutter/material.dart';
import 'package:frontend/presentation/controller/exchange_rate.dart';
import 'package:frontend/presentation/widget/custom_error.dart';

import '../../domain/entity/exchange_rate.dart';
import '../../l10n/app_localizations.dart';

class Toolbar extends StatefulWidget {
  final ExchangeRateController controller;
  const Toolbar({super.key, required this.controller});

  @override
  State<Toolbar> createState() => _ToolbarState();
}

class _ToolbarState extends State<Toolbar> {
  @override
  Widget build(BuildContext context) {
    final l10n = AppLocalizations.of(context)!;

    return Padding(
      padding: const EdgeInsets.only(top: 24.0, bottom: 16.0, left: 50, right: 50),
      child: Row(
        children: [
          ValueListenableBuilder<bool>(valueListenable:
          widget.controller.isLoading,
              builder: (context, isLoading, child) {

                if (isLoading) {
                  return const CircularProgressIndicator();
                }

                return ValueListenableBuilder<String>(
                    valueListenable: widget.controller.isError,
                    builder: (context, error, child){
                      if (error != "") {
                        return CustomError();
                      }

                      return ValueListenableBuilder<ExchangeRateEntity?>(
                          valueListenable: widget.controller.exchangeRate,
                          builder: (context, res, child){

                            if (res == null){
                              return Text(l10n.emptyRealEstateList);
                            }

                            return Container(
                              padding: const EdgeInsets.symmetric(horizontal: 12, vertical: 8),
                              decoration: BoxDecoration(
                                color: const Color(0xFFF0FDF4),
                                borderRadius: BorderRadius.circular(8),
                                border: Border.all(color: const Color(0xFFBBF7D0)),
                              ),
                              child: Row(
                                mainAxisSize: MainAxisSize.min,
                                children: [
                                  const Icon(
                                    Icons.info_outline,
                                    size: 16,
                                    color: Color(0xFF166534),
                                  ),
                                  const SizedBox(width: 8),
                                  RichText(
                                    text: TextSpan(
                                      text: l10n.quoteSubtitle,
                                      style: TextStyle(color: Color(0xFF166534), fontSize: 14),
                                      children: [
                                        TextSpan(
                                          text: 'R\$ 1,00 = \$${res.value.toStringAsFixed(2)}',
                                          style: TextStyle(fontWeight: FontWeight.bold),
                                        ),
                                      ],
                                    ),
                                  ),
                                ],
                              ),
                            );
                          }
                      );
                    });
              }
          )
        ],
      ),
    );
  }
}


