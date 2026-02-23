import 'package:flutter/material.dart';
import 'package:frontend/presentation/widget/add_property.dart';

import 'package:frontend/presentation/widget/header.dart';
import 'package:get_it/get_it.dart';

import '../controller/address.dart';
import '../controller/exchange_rate.dart';
import '../controller/real_estate.dart';
import '../widget/exchange_rate.dart';
import '../widget/real_estate.dart';
import '../widget/toast_listener.dart';
import '../widget/toolbar.dart';


class Home extends StatefulWidget {

  const Home({super.key});
  @override
  State<Home> createState() => _HomeState();
}

class _HomeState extends State<Home> {
  final realEstateController = GetIt.I<RealEstateController>();
  final exchangeRateController = GetIt.I<ExchangeRateController>();
  final addressController = GetIt.I<AddressController>();



  var nav = "home";

  void navigate(String page){

    setState(() {
      nav = page;
    });
    return;
  }

  @override
  void initState() {
    super.initState();

    exchangeRateController.getLatestExchangeRate();
  }


  @override
  Widget build(BuildContext context) {

    return Scaffold(
      appBar: AppBar(
        backgroundColor: Colors.white,
        elevation: 0,
        scrolledUnderElevation: 0,
        bottom: PreferredSize(
          preferredSize: const Size.fromHeight(1.0),
          child: Container(
            color: Colors.grey.shade200,
            height: 2.0,
          ),
        ),
        title: Header(navigate: navigate),
      ),
      body: ToastListener(
        errorNotifiers: [
          realEstateController.createIsError,
          exchangeRateController.createError,
        ],
        successNotifiers: [
          exchangeRateController.createSuccess,
          realEstateController.createSuccess
        ],
        child: Center(
          child: Column(
            mainAxisAlignment: MainAxisAlignment.center,
            children: <Widget>[
              if (nav == "home")
                Toolbar(controller: exchangeRateController),
              if(nav == "home")
               Expanded(child: RealEstateList(controller: realEstateController, exchangeRateController: exchangeRateController,)),
              if(nav == "quote")
                ExchangeRate(navigate: navigate,controller: exchangeRateController,),
              if (nav == "add")
                Expanded(child: AddProperty(navigate: navigate,realEstateController: realEstateController,addressController: addressController,)),
            ],
          ),
        ),
      ),
    );
  }
}