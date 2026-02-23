import 'package:frontend/domain/repository/exchange_rate.dart';

import '../entity/exchange_rate.dart';



class GetLatestExchangeRateUseCase{
  final ExchangeRateRepository exchangeRateRepository;

  GetLatestExchangeRateUseCase({required this.exchangeRateRepository});

  Future<ExchangeRateEntity> getLatestExchangeRate() async{
    return await exchangeRateRepository.getLatestExchangeRate();
  }

}

class AddExchangeRateUseCase{
  final ExchangeRateRepository exchangeRateRepository;

  AddExchangeRateUseCase({required this.exchangeRateRepository});

  Future<String> addExchangeRate(ExchangeRateEntity exchangeRate) async{
    return await exchangeRateRepository.addExchangeRate(exchangeRate);
  }
}