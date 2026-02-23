
import 'package:flutter/material.dart';
import 'package:frontend/domain/entity/exchange_rate.dart';
import 'package:frontend/domain/entity/real_estate.dart';
import 'package:frontend/presentation/controller/exchange_rate.dart';
import 'package:frontend/presentation/widget/pagination.dart';
import 'package:frontend/presentation/widget/property_card.dart';

import '../../domain/entity/pagination.dart';
import '../../l10n/app_localizations.dart';
import '../controller/real_estate.dart';
import 'custom_error.dart';


class RealEstateList extends StatefulWidget {
  final RealEstateController controller;
  final ExchangeRateController exchangeRateController;



  const RealEstateList({super.key, required this.controller,required this.exchangeRateController});



  @override
  State<RealEstateList> createState() => _RealEstateListState();
}

class _RealEstateListState extends State<RealEstateList> {
  double quote = 0;

  @override
  void initState() {
    super.initState();

    widget.exchangeRateController.getLatestExchangeRate();
    widget.controller.getRealEstate(1, 9);
  }

  @override
  Widget build(BuildContext context) {
    final l10n = AppLocalizations.of(context)!;

    return ValueListenableBuilder<bool>(
      valueListenable: widget.controller.isLoading,
      builder: (context, isLoading, child) {

        if (isLoading) {
          return SizedBox( width:50, height:50, child: const CircularProgressIndicator());
        }

        return ValueListenableBuilder<String>(
          valueListenable: widget.controller.isError,
          builder: (context, error, child) {

            if (error != "") {
              return CustomError();
            }

            return ValueListenableBuilder<PaginationEntity<RealEstateEntity>?>(
              valueListenable: widget.controller.realEstate,
              builder: (context, res, child) {

                if (res == null){
                  return CustomError();
                }
                if (res.data.isEmpty) {
                  return Text(l10n.emptyRealEstateList);
                }

                return Column(
                  children: [
                    Expanded(
                      child: SingleChildScrollView(
                        child: Padding(
                          padding: const EdgeInsets.all(16.0),
                          child: Wrap(
                            spacing: 16.0,
                            runSpacing: 16.0,
                            alignment: WrapAlignment.start,
                            children: res.data.map((realEstate) {
                              return ConstrainedBox(
                                constraints: const BoxConstraints(maxWidth: 350),
                                child: ValueListenableBuilder<ExchangeRateEntity?>(
                                  valueListenable: widget.exchangeRateController.exchangeRate,
                                  builder: (context, res, child) {
                                    if (res == null){
                                      return PropertyCard(property: realEstate, quote: 0);
                                    }


                                    return PropertyCard(property: realEstate, quote: res.value);
                                  }
                                )
                              );
                            }).toList(),
                          ),
                        ),
                      ),
                    ),
                    Pagination(
                      currentPage: res.page,
                      totalPages: res.totalPages,
                      onPageChanged: (newPage) {
                        widget.controller.getRealEstate(newPage, 10);
                      },
                    ),
                  ],
                );
              },
            );
          },
        );
      },
    );
  }
}
