import scrapy

from scrapy_baidu.items import ScrapyBaiduItem


class BaiduSpider(scrapy.Spider):
    name = 'baidu'
    allowed_domains = ['baidu.com']
    start_urls = ['https://www.baidu.com/s?ie=UTF-8&wd=crawlab']

    def parse(self, response):
        for el in response.css('#content_left > .result.c-container'):
            item = ScrapyBaiduItem(
                name=''.join(el.css('h3.t *::text').getall()),
                url=el.css('h3.t a::attr("href")').get(),
                abstract=''.join(el.css('.c-abstract *::text').getall()),
            )
            yield item

        next_url = response.css('#page a.n::attr("href")').get()
        if next_url is not None:
            yield scrapy.Request(
                url=f'https://www.baidu.com{next_url}'
            )
