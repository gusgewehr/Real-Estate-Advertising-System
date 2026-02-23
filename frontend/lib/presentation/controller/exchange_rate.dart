
import 'package:flutter/cupertino.dart';

import '../../domain/entity/exchange_rate.dart';
import '../../domain/usecase/exchange_rate.dart';

class ExchangeRateController{
  final GetLatestExchangeRateUseCase getLatestExchangeRateUseCase;
  final AddExchangeRateUseCase addExchangeRateUseCase;


  ExchangeRateController({required this.getLatestExchangeRateUseCase, required this.addExchangeRateUseCase});

  final ValueNotifier<bool> isLoading = ValueNotifier<bool>(false);
  final ValueNotifier<String> isError = ValueNotifier<String>("");
  final ValueNotifier<String> createSuccess = ValueNotifier<String>("");
  final ValueNotifier<ExchangeRateEntity?> exchangeRate = ValueNotifier<ExchangeRateEntity?>(null);

  Future<double> getLatestExchangeRate() async{
    isLoading.value = true;
    isError.value = "";

    try{
      final res = await getLatestExchangeRateUseCase.getLatestExchangeRate();

      exchangeRate.value = res;
      return res.value;
    } catch(e){
      isError.value = e.toString();
    } finally{
      isLoading.value = false;
    }
    return 0;
  }

  Future<void> addExchangeRate(double exchangeRate) async{
    isLoading.value = true;
    isError.value = "";

    ExchangeRateEntity exchangeRateEntity = ExchangeRateEntity(value: exchangeRate);


    try{
      await addExchangeRateUseCase.addExchangeRate(exchangeRateEntity);

      createSuccess.value = "Exchange rate updated successfully";
    }catch(e){
      isError.value = e.toString();
    } finally{
      isLoading.value = false;
      getLatestExchangeRate();
    }

  }




}