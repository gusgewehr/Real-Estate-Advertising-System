import '../entity/exchange_rate.dart';

abstract class ExchangeRateRepository{
  Future<ExchangeRateEntity> getLatestExchangeRate();
  Future<String> addExchangeRate(ExchangeRateEntity exchangeRate);
}